#include <stdio.h>
#include <string.h>
#include <lua.h>
#include <lauxlib.h>
#include <lualib.h>

int add(lua_State* L)
{
    double op1 = luaL_checknumber(L, 1);
    double op2 = luaL_checknumber(L, 2);

    lua_pushnumber(L, lua_gettop(L));
    lua_pushnumber(L, op1 + op2);

    return 2;
}

int sub(lua_State* L)
{
    double op1 = luaL_checknumber(L, 1);
    double op2 = luaL_checknumber(L, 2);

    lua_pushnumber(L, lua_gettop(L));
    lua_pushnumber(L, op1 - op2);

    return 2;
}

static const struct luaL_Reg mylibs[] = {
    {"add", add},
    {"sub", sub},
    {NULL, NULL}
};

// here mylib is a module name and must as same as the file name
int luaopen_mylib(lua_State* L)
{
    luaL_newlib(L, mylibs);
    return 1;
}
/*  compile the file command */
// gcc mytestlib.c -lm -ldl -fPIC -shared -o mylib.so
