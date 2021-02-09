do
  local old_open = io.open
  local access_ok = function (filename, mode)
    -- check access
  end
  io.open = function (filename, mode)
    if access_ok(filename, mode) then
      return old_open(filename, mode)
    else
      return nil, "access denied"
    end
  end
end
