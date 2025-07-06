#!/bin/bash

# mfp-xxx bash autocompletion script.
#
# put this file in the /etc/bash_completion.d/ or source it from your
# $HOME/.bashrc

__mfp_complete()
{
	local	cmd args words cword cur cur_raw reply r prefix

	# Shell passes us command line split into words in the
	# COMP_WORDS array, but its idea about word separator
	# characters doesn't match our goals. In particular, it
	# considers '=' character the word separator, which breaks
	# the long options and considers ':' character the word
	# separator, which breaks the host:port parameters and
	# HTTP URLs.
	#
	# _comp_get_words from the bash-completion package re-splits
	# the command line allowing some characters to be excluded
	# from the word separators list with the -n option. Nut now
	# we have another problem: shell automatically removes current
	# word prefix from the COMPREPLY strings, but after re-split the
	# current word might have been extended with the additional
	# prefix which shell doesn't know about and hence cannot
	# remove, so we have to help it.
	_comp_get_words -n "=:" words cword cur
	cur_raw="${COMP_WORDS[COMP_CWORD]}"

	if [ "${cur}" != "${cur_raw}" ]; then
		prefix="${cur}"
	fi

	cmd="$1"
	args=("${words[@]:1:$cword}")
	IFS=$'\n' read -r -d '' -a reply < <("$cmd" --bash-completion "${args[@]}" && printf '\0')

	# Build COMPREPLY, removing prefix
	COMPREPLY=()
	for r in "${reply[@]}"; do
		COMPREPLY+=("${r#"${prefix}"}")
	done
}

complete -o nospace -F __mfp_complete mfp-cups
complete -o nospace -F __mfp_complete mfp-discover
complete -o nospace -F __mfp_complete mfp-model
complete -o nospace -F __mfp_complete mfp-proxy
complete -o nospace -F __mfp_complete mfp-virtual
