function read_whole_file(file)
  local f = assert(io.open(file, "rb"))
  local content = f:read("*all")
  f:close()
  return content
end

function write_file(file, content)
  local f = assert(io.open(file, "w"))
  f:write(content)
  f:close()
end

-- 把以下划线结尾的单词替换为以 m_ 开头的单词
function rename_class_member(input_file_name, output_file_name)
  local input_content = read_whole_file(input_file_name)
  output_content = string.gsub(input_content, "(%w[%w_]+%w)_([%W])", "m_%1%2")
  write_file(output_file_name, output_content)
end

function main()
  local input_file_name = "/home/huangjian/test/test.hpp"
  local output_file_name = "/home/huangjian/test/test_out.hpp"
  rename_class_member(input_file_name, output_file_name)
end

main()
