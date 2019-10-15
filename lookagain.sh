find -name "*.sh" | sed 's/.sh/ /g' | cat -e | tr -d './'
