#!/usr/bin/tclsh

# Prompt the user to enter the number of times to print "hi"
puts  -nonewline "Enter the number of times to print 'hi': "
flush stdout
gets stdin numTimes

# Check if the input is a valid number
if {[string is integer $numTimes]} {
    for {set i 0} {$i < $numTimes} {incr i} {
        puts "hi"
    }
} else {
    puts "Please enter a valid integer."
}
