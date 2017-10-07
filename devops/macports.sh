#!/usr/bin/env bash

export LANG="C"
export LC_ALL="C" 

developer_essentials_file="macports.developer_essentials.txt"
developer_extras_file="macports.developer_extras.txt"

get_category()
{
	port="$1"
	# Example (running `port info --category coreutils`):
	# category: sysutils
	result="$(port info --category "${port}")"
	case "${result}" in
	category:\ *)
		categories="${result##category: }"
		category=${categories%%,*}
		echo "${category}"
		;;
	esac
}

get_portfile()
{
	port="$1"
	category="$(get_category "${port}")"
	if [ -n "${category}" ]
	then
		echo "https://trac.macports.org/browser/trunk/dports/${category}/${port}/Portfile"
	fi
}

print_index()
{
	ports_file="$1"
	while read port
	do
		echo "  - [${port}](#${port})"
	done <"${ports_file}"
}

print_ports()
{
	ports_file="$1"
	while read port
	do
		print_port "${port}"
	done <"${ports_file}"
}

print_port()
{
	port="$1"	
	# Example (running `port contents dos2unix`):
	# Port dos2unix contains:
	#   /opt/local/bin/dos2unix
	#   /opt/local/bin/mac2unix
	#   /opt/local/bin/unix2dos
	#   /opt/local/bin/unix2mac
	#   /opt/local/share/doc/dos2unix-7.2/BUGS.txt
	#   /opt/local/share/doc/dos2unix-7.2/COPYING.txt
	# ...
	port contents "${port}" |
	{
		is_contents=0
		echo
		while read line
		do
			if [ "$is_contents" -eq 0 ]
			then
				if [ "${line}" = "Port ${port} contains:" ]
				then
					portfile="$(get_portfile "${port}")"
					echo "##### [${port}](${portfile})"
					is_contents=1
				fi
			else
				file="${line}"

				# Check executable first
				if [ -x "${file}" ]
				then
					case "${file}" in
					*.so | \
					*.dylib)
						# 1: Unix style shared libraries
						# 2: Mac OS X style shared libraries
						#echo "L ${file}"
						;;
					/opt/local/bin/* | \
					/opt/local/libexec/gnubin/*)
						# 1: Executables
						# 2: Executables without the "g" prefix
						echo "- ${file}"
						;;
					/opt/local/libexec/*)
						# 1: Executables not designed to be called directly
						echo "- ${file}"
						;;
					/opt/local/*/bin/*)
						# 1: Executables hidden to avoid name conflicts
						echo "- ${file}"
						;;
					*)
						echo "${file} <-- Unknown" >&2
					esac
				fi

				# Check path first
				case "${file}" in
				*.a | \
				*.h | \
				*.awk | \
				/opt/local/share/doc/* | \
				/opt/local/share/info/* | \
				/opt/local/share/locale/*/LC_MESSAGES/*.mo | \
				/opt/local/share/locale/*/LC_TIME/*.mo | \
				/opt/local/share/man/man?/* | \
				/opt/local/share/man/*/man?/* | \
				/opt/local/libexec/gnubin/man/man?/*)
					if [ -x "${file}" ]
					then
						echo "${file} <-- Wrong permission" >&2
					fi
					;;
				*.so | \
				*.dylib | \
				/opt/local/bin/* | \
				/opt/local/libexec/gnubin/* | \
				/opt/local/libexec/* | \
				/opt/local/*/bin/*)
					if [ ! -x "${file}" ]
					then
						echo "${file} <-- Wrong permission" >&2
					fi
					;;
				*)
					echo "${file} <-- Unknown (2)" >&2
				esac
			fi
		done
	}
}

echo '# MacPorts Packages'
echo
echo '### Index'
echo
echo '- Developer Essentials'
print_index "${developer_essentials_file}"
echo '- Developer Extras'
print_index "${developer_extras_file}"
echo
echo '### Developer Essentials'
print_ports "${developer_essentials_file}"
echo
echo '### Developer Extras'
print_ports "${developer_extras_file}"
