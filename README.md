# mkpasswd

A command line tool for generating random passwords.  
Life is hard, just smile. :smile:

## Features

- Command line with parameters
- Support for generating multiple passwords
- Arabic numerals supported `0-9`
- Special characters supported `!@#$%^&*+?`
- `Makefile` is supported

## Usage

The command includes three optional parameters.

```sh
mkpasswd -N 5 -l 32 -n 8 -c 8
```

- `-N`: The quantity of created passwords.
- `-l`: The length of password.
- `-n`: The number of Arabic numerals in password.
- `-c`: The number of special characters in password.

## License

GPL-3.0 License
