# patterns

```lua
.    -- all characters
%a   -- letters
%c   -- control characters
%d   -- digits
%g   -- printable characters except spaces
%l   -- lower-case letters
%p   -- punctuation characters
%s   -- space characters
%u   -- upper-case letters
%w   -- alphanumeric characters
%x   -- hexadecimal digits

-- 大写字母代表相反的意思，例如：
%A   -- represents all non-letter characters
```

```lua
-- magic characters
( ) . % + - * ? [ ] ^ $
```

### gsub

http://www.lua.org/manual/5.3/manual.html#pdf-string.gsub

```txt
The character % works as an escape character: any sequence in repl of the form
%d, with d between 1 and 9, stands for the value of the d-th captured substring.
The sequence %0 stands for the whole match. The sequence %% stands for a single
%.
```

