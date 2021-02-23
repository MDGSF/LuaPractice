#include <stdio.h>
#include <string.h>

extern "C" {
#include "lauxlib.h"
#include "lua.h"
#include "lualib.h"
}

int main() {
  char buff[256] = {0};
  lua_State *L = luaL_newstate();  // opens Lua
  luaL_openlibs(L);                // opens the standard libraries

  while (fgets(buff, sizeof(buff), stdin) != NULL) {
    int error = luaL_loadstring(L, buff) || lua_pcall(L, 0, 0, 0);
    if (error) {
      fprintf(stderr, "%s\n", lua_tostring(L, -1));
      lua_pop(L, 1);  // pop the error message from the stack
    }
  }

  lua_close(L);
  return 0;
}
