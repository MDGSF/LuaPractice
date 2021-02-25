#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "lua.hpp"

void error(lua_State *L, const char *fmt, ...) {
  va_list argp;
  va_start(argp, fmt);
  vfprintf(stderr, fmt, argp);
  va_end(argp);
  lua_close(L);
  exit(EXIT_FAILURE);
}

void stack_dump(lua_State *L) {
  int top = lua_gettop(L);
  printf("stack (size=%d): ", top);
  for (int i = 1; i <= top; i++) {
    printf("[");
    int t = lua_type(L, i);
    switch (t) {
      case LUA_TSTRING: {
        printf("'%s'", lua_tostring(L, i));
        break;
      }
      case LUA_TBOOLEAN: {
        printf(lua_toboolean(L, i) ? "true" : "false");
        break;
      }
      case LUA_TNUMBER: {
        if (lua_isinteger(L, i)) {
          printf("%lld", lua_tointeger(L, i));
        } else {
          printf("%g", lua_tonumber(L, i));
        }
        break;
      }
      default: {
        printf("%s", lua_typename(L, t));
        break;
      }
    }
    printf("] ");
  }
  printf("\n");
}