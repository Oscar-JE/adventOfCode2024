{:ok, input} = File.read("example.txt")
IO.puts(input)
splited = String.split(input, "\n")
IO.puts(splited)
l1 = []
l2 = []
for el
