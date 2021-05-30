def isValid(s)
    freq = Hash.new

    s.split('').each do |c|
        if freq[c] == nil
            freq[c] = 1
        else 
            freq[c] = freq[c] + 1
        end
    end

    if freq.values.uniq.length == 1 
        print 'YES'
    else
        sorted = freq.values.sort.reverse
        sorted[0] = sorted[0] - 1
        if sorted.uniq.length == 1 
            print 'YES'
        else
            sorted = freq.values.sort.reverse
            sorted[sorted.length-1] = sorted[sorted.length-1] - 1
            if sorted[sorted.length-1] == 0
                sorted.pop
            end
            if sorted.uniq.length == 1 
                print 'YES'
            else
                print 'NO'
            end
        end
    end

    

end

isValid('aabbc')