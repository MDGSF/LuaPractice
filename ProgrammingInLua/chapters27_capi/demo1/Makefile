LUAINC=../../../LuaSourceCode/lua-5.3.6/src
LUALIB=/home/huangjian/local/lua/Binary/ubuntu16.04/lua-5.3.6/lib

name=demo1

$(name):
	g++ -Wl,--no-as-needed -ldl main.cpp -I$(LUAINC) $(LUALIB)/liblua.a -o $@

clean:
	-rm -rf $(name)

