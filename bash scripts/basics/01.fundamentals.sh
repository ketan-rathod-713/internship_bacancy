# Fundamental concepts

var=10 # assigns value to a variable

echo var # prints the string value as it is, it will not print variable value, it takes it as a string
echo $var # prints value of the variable
echo "$var" # prints the string, inside it is taking the value of the variable

echo '$var' # single quoutes will print only the string // IMPORTANT

echo {2..9} | grep "1" > /dev/null && echo "Number 1 found" || echo "Number 1 not found"
echo {1..9} | grep "1" > /dev/null && echo "Number 1 found" || echo "Number 1 not found"

# same thing {a..b} can be done with the characters.
# this syntax can also be used with touch, or other commands.

# assigning output of a command to a variable
users=$(cat /etc/passwd | awk -F":" '{print $1}')

echo $users

for i in $users; 
do echo $i; 
done

# positional parameters
echo $0 # script name
echo $1 # first argument
echo $2 # second argument
echo $# # number of arguments
echo $@ # all arguments
echo $* # all arguments, space separated

# if arguments are not provided, then empty string is assigned