// MFP - Miulti-Function Printers and scanners toolkit
// IPP registrations to Go converter.
//
// Copyright (C) 2024 and up by Alexander Pevzner (pzz@apevzner.com)
// See LICENSE for license terms and conditions
//
// Registrations database

package main

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/OpenPrinting/go-mfp/util/generic"
	"github.com/OpenPrinting/go-mfp/util/xmldoc"
)

// RegDB represents database of IANA registrations for IPP
type RegDB struct {
	Collections map[string]map[string]*RegDBAttr // Attrs by collection
	AllAttrs    map[string]*RegDBAttr            // All attributes, by path
	Links       map[string]string                // Src uses target's attrs
	ErrataSkip  generic.Set[string]              // Errata: ignored attrs
	Errata      map[string]*RegDBAttr            // Errata: replaced by path
	Errors      []error                          // Collected errors
}

// RegDBAttr represents a single attribute
type RegDBAttr struct {
	Name         string                // Attribute name
	Collection   string                // Collection it belongs to
	Parents      []string              // Parent collections
	SyntaxString string                // Syntax string
	Syntax       Syntax                // Attribute syntax, parsed
	XRef         string                // Document it is defined in
	Members      map[string]*RegDBAttr // Members, by name
}

// NewRegDB creates a new RegDB
func NewRegDB() *RegDB {
	return &RegDB{
		Collections: make(map[string]map[string]*RegDBAttr),
		AllAttrs:    make(map[string]*RegDBAttr),
		Links:       make(map[string]string),
		ErrataSkip:  generic.NewSet[string](),
		Errata:      make(map[string]*RegDBAttr),
	}
}

// Load loads attributes from XML
func (db *RegDB) Load(xml xmldoc.Element, errata bool) error {
	for _, registry := range xml.Children {
		// Ignore elements other that "registry"
		if registry.Name != "registry" {
			continue
		}

		// Process "record" elements
		for _, record := range registry.Children {
			// Ignore elements other that "record"
			if record.Name != "record" {
				continue
			}

			// Lookup fields we are interested in
			collection := xmldoc.Lookup{Name: "collection", Required: true}
			name := xmldoc.Lookup{Name: "name", Required: true}
			syntax := xmldoc.Lookup{Name: "syntax", Required: true}
			xref := xmldoc.Lookup{Name: "xref", Required: true}
			member := xmldoc.Lookup{Name: "member_attribute"}
			submember := xmldoc.Lookup{Name: "sub-member_attribute"}

			missed := record.Lookup(&collection, &name, &syntax, &xref,
				&member, &submember)

			// Ignore elements with missed fields
			if missed != nil {
				continue
			}

			// Elements with empty syntax must be links
			if syntax.Elem.Text == "" {
				from, to, err := db.newLink(
					collection.Elem.Text,
					name.Elem.Text,
					member.Elem.Text,
					submember.Elem.Text,
				)
				if err == nil && !db.ErrataSkip.Contains(from) {
					err = db.newDirectLink(from, to)
				}
				if err != nil {
					return err
				}
				continue
			}

			// Create attribute
			attr, err := db.newRegDBAttr(
				collection.Elem.Text,
				name.Elem.Text,
				member.Elem.Text,
				submember.Elem.Text,
				syntax.Elem.Text,
				xref.Elem.Text,
			)

			if err != nil {
				return err
			}

			// Add to the database
			if errata {
				err = db.addErrata(attr)
			} else {
				if !db.ErrataSkip.Contains(attr.Path()) {
					err = db.add(attr)
				}
			}

			if err != nil {
				return err
			}
		}

		// Process "link" and "skip" elements; only in errata
		if errata {
			for _, chld := range registry.Children {
				switch chld.Name {
				case "link":
					from := xmldoc.Lookup{Name: "from", Required: true}
					to := xmldoc.Lookup{Name: "to", Required: true}

					missed := chld.Lookup(&from, &to)
					if missed != nil {
						err := fmt.Errorf("missed link element %s", missed.Name)
						return err
					}

					// Create the link
					err := db.newDirectLink(from.Elem.Text, to.Elem.Text)
					if err != nil {
						return err
					}

				case "skip":
					name, ok := chld.ChildByName("name")
					if ok {
						db.ErrataSkip.Add(name.Text)
					}
				}
			}
		}
	}

	if !errata {
		db.verifyLinks()
		db.handleSuffixes()
		db.checkEmptyCollections()
	}

	return nil
}

// addErrata adds attribute to the db.Errata
func (db *RegDB) addErrata(attr *RegDBAttr) error {
	path := attr.Path()

	if db.Errata[path] != nil {
		err := fmt.Errorf("%s: duplicated errata attribute", path)
		return err
	}

	db.Errata[path] = attr
	return nil
}

// add adds attribute to the database
func (db *RegDB) add(attr *RegDBAttr) error {
	// Make collection on demand
	collection := db.Collections[attr.Collection]
	if collection == nil {
		collection = make(map[string]*RegDBAttr)
		db.Collections[attr.Collection] = collection
	}

	// Check parents and build path
	var parent *RegDBAttr
	path := []string{attr.Collection}

	for i := range attr.Parents {
		path = append(path, attr.Parents[i])
		if parent != nil {
			parent = parent.Members[attr.Parents[i]]
		} else {
			parent = collection[attr.Parents[i]]
		}

		if parent == nil {
			err := fmt.Errorf("%s: no parent (%s)",
				attr.Path(), strings.Join(path, "/"))
			return err
		}
	}

	// Check for duplicates and save attribute
	pmap := collection
	if parent != nil {
		pmap = parent.Members
	}

	if attr2 := pmap[attr.Name]; attr2 != nil {
		err := fmt.Errorf("%s: conflicts with %s",
			attr.Path(), attr2.Path())
		return err
	}

	db.AllAttrs[attr.Path()] = attr
	pmap[attr.Name] = attr
	return nil
}

// CollectionNames returns names of attribute collections
func (db *RegDB) CollectionNames() []string {
	collections := make([]string, 0, len(db.Collections))
	for col := range db.Collections {
		collections = append(collections, col)
	}

	sort.Strings(collections)
	return collections
}

// Lookup returns attributes, associated with path.
// Path could look as follows:
//
//	Description/input-attributes-default - path to particular attribute
//	Job Template                         - path to entire collection
func (db *RegDB) Lookup(path string) map[string]*RegDBAttr {
	if i := strings.IndexByte(path, '/'); i < 0 {
		// Top-level collection requested
		return db.Collections[path]
	}

	// Particular attribute requested
	attr := db.AllAttrs[path]
	if attr != nil {
		return attr.Members
	}

	return nil
}

// verifyLinks checks that there is no broken links in the db.Links
func (db *RegDB) verifyLinks() {
	// Process links in predictable way
	sources := make([]string, 0, len(db.Links))
	for src := range db.Links {
		sources = append(sources, src)
	}

	sort.Strings(sources)

	// Now verify
	for _, src := range sources {
		dst := db.Links[src]

		if db.Lookup(src) == nil {
			err := fmt.Errorf("%s->%s: broken source", src, dst)
			db.Errors = append(db.Errors, err)

		} else if db.Lookup(dst) == nil {
			err := fmt.Errorf("%s->%s: broken target", src, dst)
			db.Errors = append(db.Errors, err)
		}
	}
}

// handleSuffixes handles attributes, marked by suffixes
// ("(extension)", "(deprecated)" etc) in their names.
func (db *RegDB) handleSuffixes() {
	collections := db.CollectionNames()
	for _, col := range collections {
		db.handleSuffixesRecursive(db.Collections[col])
	}
}

// handleSuffixesRecursive does the real work of handling
// attribute suffixes.
func (db *RegDB) handleSuffixesRecursive(attrs map[string]*RegDBAttr) {
	// Gather aliases.
	//
	// Aliases are attributes with the same base name, but different
	// suffixes.
	allAliases := make(map[string][]*RegDBAttr)

	for _, attr := range attrs {
		name := attr.PureName()
		aliases := allAliases[name]
		aliases = append(aliases, attr)
		allAliases[name] = aliases
	}

	// Prepare sorted list of names, so things will be done
	// in the reproducible way.
	names := make([]string, 0, len(allAliases))
	for name := range allAliases {
		names = append(names, name)
	}

	sort.Strings(names)

	// Resolve aliases and rebuild attrs
	clear(attrs)
	for _, name := range names {
		aliases := allAliases[name]
		attr := aliases[0]
		if len(aliases) > 0 {
			var err error
			attr, err = db.resolveAliases(aliases)
			if err != nil {
				db.Errors = append(db.Errors, err)
			}
		}

		if attr != nil {
			attrs[name] = attr
		}
	}

	for _, name := range names {
		attr := attrs[name]
		if attr != nil {
			db.handleSuffixesRecursive(attr.Members)
		}
	}
}

// checkEmptyCollections checks for collection attributes
// without members.
func (db *RegDB) checkEmptyCollections() {
	collections := db.CollectionNames()
	for _, col := range collections {
		db.checkEmptyCollectionsRecursive(db.Collections[col])
	}
}

// checkEmptyCollectionsRecursive does the real work of checking
// for empty collections.
func (db *RegDB) checkEmptyCollectionsRecursive(attrs map[string]*RegDBAttr) {
	// Prepare sorted list of names, so things will be done
	// in the reproducible way.
	names := make([]string, 0, len(attrs))
	for name := range attrs {
		names = append(names, name)
	}

	sort.Strings(names)

	// Now roll over all names
	for _, name := range names {
		attr := attrs[name]
		if attr.Syntax.Collection && len(attr.Members) == 0 && db.Links[attr.Path()] == "" {
			err := fmt.Errorf("%s: empty collection", attr.Path())
			db.Errors = append(db.Errors, err)
		}
	}
}

func (db *RegDB) resolveAliases(aliases []*RegDBAttr) (*RegDBAttr, error) {
	var plain *RegDBAttr
	deprecated := []*RegDBAttr{}
	extension := []*RegDBAttr{}

	// Classify aliases
	for _, attr := range aliases {
		_, suffixes := attr.SplitName()
		switch {
		case suffixes == "":
			plain = attr
		case strings.Contains(suffixes, "extension"):
			extension = append(extension, attr)
		default:
			deprecated = append(deprecated, attr)
		}
	}

	// Now choose candidates
	candidates := extension
	if len(candidates) == 0 && plain != nil {
		candidates = append(candidates, plain)
	}
	if len(candidates) == 0 {
		candidates = append(candidates, deprecated...)
	}

	// Merge equal candidates
	end := 1
	for i := 1; i < len(candidates); i++ {
		attr := candidates[i]
		prev := candidates[i-1]

		if !attr.Syntax.Equal(prev.Syntax) {
			candidates[end] = attr
			end++
		}
	}

	candidates = candidates[:end]

	// If we only have a single candidate, everything is OK
	if len(candidates) == 1 {
		return candidates[0], nil
	}

	// Format the error message
	buf := &bytes.Buffer{}
	fmt.Fprintf(buf, "Conflicting attribytes:\n")
	for _, attr := range candidates {
		fmt.Fprintf(buf, "  %s\n", attr.Path())
		fmt.Fprintf(buf, "  > %s\n", attr.SyntaxString)
	}

	return nil, errors.New(buf.String())
}

// newLink prepares new link from one collection attribute to another.
// Link means that one attribute borrows members from another.
//
// In the XML it looks as follows:
//
//	<record>
//	  <collection>Document Status</collection>
//	  <name>cover-back-actual</name>
//	  <member_attribute>&lt;Any "cover-back" member attribute&gt;</member_attribute>
//	  <syntax/>
//	  <xref type="uri" data="https://ftp.pwg.org/pub/pwg/candidates/cs-ippdocobject12-20240517-5100.5.pdf">PWG5100.5</xref>
//	</record>
//
// On success, additional explicit call to RegDB.newDirectLink is required to
// save the link.
func (db *RegDB) newLink(collection, name, member, submember string) (
	from, to string, err error) {

	var path []string
	var link []string
	var last string

	switch {
	case name != "" && member != "" && submember == "":
		path = []string{collection, name}
		link = []string{collection}
		last = member

	case name != "" && member != "" && submember != "":
		path = []string{collection, name, member}
		link = []string{collection, name}
		last = submember

	default:
		panic("internal error")
	}

	if fields := strings.Split(last, `"`); len(fields) == 3 {
		// <Any "cover-back" member attribute>
		// <-00> <----1---> <--------2------->
		last = fields[1]
		link = append(link, last)
	} else {
		// <Any Job Template attribute>
		last = strings.TrimPrefix(last, "<Any ")
		last = strings.TrimSuffix(last, " attribute>")
		last = strings.TrimSuffix(last, " Attribute>")
		link = []string{last}
	}

	from = strings.Join(path, "/")
	to = strings.Join(link, "/")

	return
}

// newDirectLink creates new link directly, using absolute paths.
// It allows to create link targeted to the different top-level collection.
// Used only in the errata.xml with the following syntax:
func (db *RegDB) newDirectLink(from, to string) error {
	if from == to {
		// Sometimes we have link to self, just skip it
		return nil
	}

	if db.Links[from] != "" {
		err := fmt.Errorf("%s: duplicated link", from)
		return err
	}

	//fmt.Println(from, "->", to)
	db.Links[from] = to

	return nil
}

// newRegDBAttr creates a new RegDBAttr
func (db *RegDB) newRegDBAttr(collection, name, member, submember,
	syntax, xref string) (*RegDBAttr, error) {

	// Create RegDBAttr structure
	attr := &RegDBAttr{
		Collection:   collection,
		SyntaxString: syntax,
		XRef:         xref,
		Members:      make(map[string]*RegDBAttr),
	}

	// Populate Name and Parents
	switch {
	case name != "" && member == "" && submember == "":
		attr.Name = name

	case name != "" && member != "" && submember == "":
		attr.Name = member
		attr.Parents = []string{name}

	case name != "" && member != "" && submember != "":
		attr.Name = submember
		attr.Parents = []string{name, member}

	default:
		panic("internal error")
	}

	// Check for errata. Do it before we attempt to
	// parse the syntax.
	if errata := db.Errata[attr.Path()]; errata != nil {
		return errata, nil
	}

	// Parse syntax
	var err error
	attr.Syntax, err = ParseSyntax(syntax)
	if err != nil {
		err = fmt.Errorf("%s: %w", attr.Path(), err)
	}

	return attr, err
}

// SplitName returns attribute's base name and suffixes, separated.
func (attr *RegDBAttr) SplitName() (name, suffixes string) {
	if i := strings.IndexByte(attr.Name, '('); i >= 0 {
		return attr.Name[:i], attr.Name[i:]
	}

	return attr.Name, ""
}

// PureName returns attribute name with possible suffixes stripped
func (attr *RegDBAttr) PureName() string {
	name, _ := attr.SplitName()
	return name
}

// Path returns full path to the attribute in the following form:
// "Collection/name/member/submember"
func (attr *RegDBAttr) Path() string {
	path := []string{attr.Collection}
	path = append(path, attr.Parents...)
	path = append(path, attr.Name)
	return strings.Join(path, "/")
}
