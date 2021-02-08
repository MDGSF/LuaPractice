function BubbleSort (a)
  for i = 1, #a do
    for j = 1, #a - i do
      if a[j] > a[j+1] then
        a[j], a[j+1] = a[j+1], a[j]
      end
    end
  end
end

function ShowTables (a)
  for k, v in ipairs(a) do
    print(k, v)
  end
end

a = { 2, 1, 3, 6, 4, 5 }
ShowTables(a)

BubbleSort(a)

print("-------------")
ShowTables(a)
