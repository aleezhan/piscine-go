curl -s https://api.github.com/users/aleezhan | grep ' "id"' | cut -d : -f 2 | cut -d , -f1 | cut -d " " -f2
