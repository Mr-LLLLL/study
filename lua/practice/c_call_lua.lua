str = "my lua lib"

person = {name = "cxl", age =26}

function add(x, y)
    print "helo"
    return x + y
end

function addName(person, newName)
    person.name = newName
    return person
end

function display()
    for k, v in pairs(person) do
        print(k .. ":   " .. v)
    end
end
