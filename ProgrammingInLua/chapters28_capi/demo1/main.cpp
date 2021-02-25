#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include "clua.h"

#define MAX_COLOR 255

struct ColorTable {
  const char *name;
  unsigned char red, green, blue;
} colortable[] = {
    {"WHITE", MAX_COLOR, MAX_COLOR, MAX_COLOR},
    {"RED", MAX_COLOR, 0, 0},
    {"GREEN", 0, MAX_COLOR, 0},
    {"BLUE", 0, 0, MAX_COLOR},
    {NULL, 0, 0, 0} /* sentinel */
};

int getglobint(lua_State *L, const char *var) {
  int isnum, result;
  lua_getglobal(L, var);  // 把 Lua 中的全局数据拷贝到虚拟栈。
  result = (int)lua_tointegerx(L, -1, &isnum);
  if (!isnum) {
    error(L, "'%s' should be a number\n", var);
  }
  lua_pop(L, 1); /* remove result from the stack */
  return result;
}

/* assume that table is on the top of the stack */
int getcolorfield(lua_State *L, const char *key) {
  int result, isnum;
  lua_pushstring(L, key); /* push key */
  lua_gettable(L, -2);    /* get background[key] */
  result = (int)(lua_tonumberx(L, -1, &isnum) * MAX_COLOR);
  if (!isnum) {
    error(L, "invalid component '%s' in color", key);
  }
  lua_pop(L, 1); /* remove number */
  return result;
}

/* assume that table is on top */
void setcolorfield(lua_State *L, const char *index, int value) {
  lua_pushstring(L, index);                     /* key */
  lua_pushnumber(L, (double)value / MAX_COLOR); /* value */
  lua_settable(L, -3);
}

void setcolor(lua_State *L, struct ColorTable *ct) {
  lua_newtable(L); /* creates a table */
  setcolorfield(L, "red", ct->red);
  setcolorfield(L, "green", ct->green);
  setcolorfield(L, "blue", ct->blue);
  lua_setglobal(L, ct->name); /* 'name' = table */
}

void get_table_color(lua_State *L, struct ColorTable *ct) {
  lua_getglobal(L, ct->name);
  if (!lua_istable(L, -1)) {
    error(L, "'%s' is not a table", ct->name);
  }
  ct->red = getcolorfield(L, "red");
  ct->green = getcolorfield(L, "green");
  ct->blue = getcolorfield(L, "blue");
}

void getcolor(lua_State *L, struct ColorTable *ct) {
  lua_getglobal(L, ct->name);
  if (lua_isstring(L, -1)) {                /* value is a string? */
    const char *name = lua_tostring(L, -1); /* get string */
    int i;                                  /* search the color table */
    for (i = 0; colortable[i].name != NULL; i++) {
      if (strcmp(name, colortable[i].name) == 0) {
        break;
      }
    }
    if (colortable[i].name == NULL) { /* string not found? */
      error(L, "invalid color name (%s)", name);
    } else { /* use colortable[i] */
      ct->red = colortable[i].red;
      ct->green = colortable[i].green;
      ct->blue = colortable[i].blue;
    }
  } else if (lua_istable(L, -1)) {
    ct->red = getcolorfield(L, "red");
    ct->green = getcolorfield(L, "green");
    ct->blue = getcolorfield(L, "blue");
  } else {
    error(L, "invalid value for '%s'", ct->name);
  }
}

void load(lua_State *L, const char *fname) {
  if (luaL_loadfile(L, fname) || lua_pcall(L, 0, 0, 0)) {
    error(L, "cannot run config. file: %s", lua_tostring(L, -1));
  }
}

int main() {
  char buff[256] = {0};
  lua_State *L = luaL_newstate();  // opens Lua
  luaL_openlibs(L);                // opens the standard libraries

  int i = 0;
  while (colortable[i].name != NULL) {
    setcolor(L, &colortable[i++]);
  }

  load(L, "conf.lua");

  int width = getglobint(L, "width");
  int height = getglobint(L, "height");
  printf("width = %d\n", width);
  printf("height = %d\n", height);

  ColorTable background;
  background.name = "background";
  getcolor(L, &background);
  printf("%s: [r:%d, g:%d, b:%d]\n", background.name, background.red,
         background.green, background.blue);

  lua_close(L);
  return 0;
}
