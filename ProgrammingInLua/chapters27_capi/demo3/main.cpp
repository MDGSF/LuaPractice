#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "lua.hpp"

int foo(lua_State *L) { return 0; }

int secure_foo(lua_State *L) {
  lua_pushcfunction(L, foo);
  return (lua_pcall(L, 0, 0, 0) == 0);
}

int main() {
  lua_State *L = luaL_newstate();  // opens Lua

  lua_close(L);
  return 0;
}
