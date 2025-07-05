#!/bin/bash

# mfp-xxx bash autocompletion script.
#
# put this file in the /etc/bash_completion.d/ or source it from your
# $HOME/.bashrc

__mfp_complete()
{
	local	cmd args

	cmd="$1"
	args=${COMP_WORDS[@]:1:$COMP_CWORD}
	COMPREPLY=()
	IFS=$'\n' read -r -d '' -a COMPREPLY < <("$cmd" --bash-completion ${args[@]} && printf '\0')
}

complete -o nospace -F __mfp_complete mfp-cups
complete -o nospace -F __mfp_complete mfp-discover
complete -o nospace -F __mfp_complete mfp-masq
complete -o nospace -F __mfp_complete mfp-model
complete -o nospace -F __mfp_complete mfp-virtual
