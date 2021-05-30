def stringConstruction(s)
    return s.chars.to_a.uniq.length
end

print stringConstruction('abcabc')