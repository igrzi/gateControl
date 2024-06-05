#!/bin/bash
 
# Acesso exec
gnome-terminal -- bash -c "
cd acesso;
export PATH=\$PATH:/usr/local/go/bin;
export GOPATH=\$HOME/go;
export PATH=\$PATH:\$GOPATH/bin;
CompileDaemon -command="./acesso";
exec bash"

# Cadastro exec
gnome-terminal -- bash -c "
cd cadastro;
export PATH=\$PATH:/usr/local/go/bin;
export GOPATH=\$HOME/go;
export PATH=\$PATH:\$GOPATH/bin;
CompileDaemon -command="./cadastro";
exec bash"

# Cancela exec
gnome-terminal -- bash -c "
cd cancela;
export PATH=\$PATH:/usr/local/go/bin;
export GOPATH=\$HOME/go;
export PATH=\$PATH:\$GOPATH/bin;
CompileDaemon -command="./cancela";
exec bash"

# Creditos exec
gnome-terminal -- bash -c "
cd creditos;
export PATH=\$PATH:/usr/local/go/bin;
export GOPATH=\$HOME/go;
export PATH=\$PATH:\$GOPATH/bin;
CompileDaemon -command="./creditos";
exec bash"

# Vagas exec
gnome-terminal -- bash -c "
cd vagas;
export PATH=\$PATH:/usr/local/go/bin;
export GOPATH=\$HOME/go;
export PATH=\$PATH:\$GOPATH/bin;
CompileDaemon -command="./vagas";
exec bash"