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

int main() {
  lua_State *L = luaL_newstate();  // opens Lua

  lua_pushboolean(L, 1);
  lua_pushnumber(L, 10);
  lua_pushnil(L);
  lua_pushstring(L, "hello");
  stack_dump(L);

  lua_pushvalue(L, -4); // 把-4下标的元素复制插入到栈顶
  stack_dump(L);

  lua_replace(L, 3); //把栈顶元素移动到3下标的位置
  stack_dump(L);

  lua_settop(L, 6);
  stack_dump(L);

  lua_rotate(L, 3, 1);
  stack_dump(L);

  lua_remove(L, -3); // 删除-3下标的元素
  stack_dump(L);

  lua_settop(L, -5);
  stack_dump(L);

  lua_close(L);
  return 0;
}
