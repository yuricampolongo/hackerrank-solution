#!/bin/ruby

require 'json'
require 'stringio'
require 'set'

#
# Complete the 'twoStrings' function below.
#
# The function is expected to return a STRING.
# The function accepts following parameters:
#  1. STRING s1
#  2. STRING s2
#

def twoStrings(s1, s2)
    first_word_set = Set.new
    second_word_set = Set.new

    s1.split('').each do |c| 
        first_word_set << c
    end

    s2.split('').each do |c| 
        second_word_set << c
    end

    intersection = first_word_set & second_word_set
    if intersection.length > 0 
        return 'YES'
    end
    return 'NO'
end

print twoStrings("aaaaaaaaaaa","bbbbbbbbbb")