a = 256

x = (1 + a) / 2.0

loop do
  ox = x
  x = (x + a.to_f / x) / 2.0
  break if x >= ox
end

puts "The square root of #{a} is #{x}"
