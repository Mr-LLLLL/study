#include <stdio.h>
#include "lua.h"
#include "lauxlib.h"
#include "lualib.h"


lua_State* load_lua(char *luaPath) 
{
    lua_State *L = luaL_newstate();
    luaL_openlibs(L);

    if (luaL_loadfile(L, luaPath) || lua_pcall(L, 0, 0, 0)) {
        printf("Load lua failed: %s\n", lua_tostring(L, -1));

        return NULL;
    }
    return L;
}


int main(int argc, char** argv)
{
    char *luaFile = "c_call_lua.lua";
    lua_State *L = load_lua(luaFile);
    if (NULL == L) {
        return -1;
    }

    printf("stack size: %d\n", lua_gettop(L));
    lua_getglobal(L, "str");
    printf("%s\n", luaL_checklstring(L, -1, NULL));

    lua_getglobal(L, "person");
    lua_getfield(L, -1, "name");
    printf("name: %s\n", luaL_checklstring(L, -1, NULL));
    lua_getfield(L, -2, "age");
    printf("age: %d\n", luaL_checkinteger(L, -1));

    lua_getglobal(L, "add");
    lua_pushnumber(L, 3);
    lua_pushnumber(L, 5);
    printf("stack size: %d\n", lua_gettop(L));
    if (lua_pcall(L, 2, 1, 0)) {
        printf("lua call failed: %s\n", luaL_checklstring(L, -1, NULL));
        return -1;
    }

    int result = luaL_checkinteger(L, -1);
    printf("result: %d\n", result);
    printf("stack size: %d\n", lua_gettop(L));
    
    lua_pushnumber(L, 165);
    lua_setfield(L, 2, "height");
    printf("stack size: %d\n", lua_gettop(L));

    /* lua_getglobal(L, "display"); */
    /* if (lua_pcall(L, 0, 0, 0)) { */
    /*     printf("lua call failed: %s\n", luaL_checklstring(L, -1, NULL)); */
    /*     return -1; */
    /* } */

    lua_getglobal(L, "addName");
    lua_newtable(L);
    lua_pushstring(L, "mzh");
    if (lua_pcall(L, 2, 1, 0)) {
        printf("lua call failed: %s\n", luaL_checklstring(L, -1, NULL));
        return -1;
    }
    lua_getfield(L, -1, "name");
    printf("name: %s\n", luaL_checklstring(L, -1, NULL));

    lua_close(L);

    return 0;
}

