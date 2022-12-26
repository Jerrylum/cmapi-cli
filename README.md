# CMAPI CLI

Command Line Interface for managing CMAPI projects. It is designed to be used with CMAPI, a C++ library for the VEX V5 used by the CMAss Robotics Team, but it can be used with any other VEX V5 project using PROS with Git version control system.

## Why CMAPI-CLI?

CMAPI-CLI is designed to speed up the development process, allowing beginners and advanced users to create, build, upload, and manage projects much more accessible. It provides a set of commands for building and managing the project by executing Git and PROS CLI commands in the correct order and with the proper parameters.

For example, the `normal` command will execute `pros make` and `pros upload` in the correct order. It first builds the project, and wait for the user to connect the V5 Brain to the computer. Then, it uploads the project to the V5 Brain. Once the upload is finished, it will sound a beep to notify the user.

On the one hand, with VS Code Pros extension, the user has to open the terminal, click `Build`, wait for the build to finish, connect the V5 Brain to the computer, click `Upload`, wait for the upload to finish, and then disconnect the V5 Brain from the computer. This process is very tedious and time-consuming. The user has to focus on the editor and wait for the build and upload to finish.

On the other hand, with CMAPI-CLI, after the user type `normal`, the project will be built and uploaded to the V5 Brain when the V5 Brain is available. They don't need to connect the V5 Brain in order. The user can focus on other things while the project is being built and will be notified when the project is built and uploaded.

Last but not least, the `normal` command is the default command, which can be executed by simply pressing enter. CMAPI-CLI is designed for professional teams to speed up their development process exponentially. It makes it easier for beginners to develop on VEX V5 too.

## Project Management

It also provides a set of commands for managing the project. All daily tasks can be done with a single command, such as `clone` to clone the repository from the Git server, then initialize the PROS project, `create` to create a new PROS project, commit the code, then create a new repository on the Git server, and `backup` to commit the changes to the project and push the changes to the remote repository.


## Get Started

To get started, please follow the instructions:

### 1. Install Git

https://git-scm.com/downloads

### 2. Install VS Code

https://code.visualstudio.com/download

### 3. Install PROS Extension

Open VS Code, and install the PROS extension.

```
Name: PROS
Id: sigbots.pros
Description: PROS Extension that allows for C/C++ Development for VEX V5 and VEX Cortex
Version: 0.4.2
Publisher: Purdue ACM SIGBots
VS Marketplace Link: https://marketplace.visualstudio.com/items?itemName=sigbots.pros
```

### 4. Install PROS-CLI and Toolchain

Click `Install PROS` in the PROS tab.

### 5. Install CMAPI-CLI

First, download the most recent build for your OS from the releases page.

https://github.com/Jerrylum/cmapi-cli/releases/latest

If you are on Windows, after you download the latest version, you need to create a folder named `cmapi-application` in your home directory. Then, move the downloaded file to the `cmapi-application` folder and rename it to `cmapi-cli.exe`.

If you are on MacOS or Linux, you can follow the instructions below to install the latest version.

```bash
# Download the latest version into the current directory
wget https://github.com/Jerrylum/cmapi-cli/releases/download/XXX/cmapi-cli-YYY

mkdir ~/cmapi-application

chmod +x cmapi-cli-*

mv cmapi-cli-* ~/cmapi-application/cmapi-cli
```

### 6. Put Everything Together

The PROS-CLI and toolchain should be installed at the following location by the extension. You need to add the following paths to your PATH environment variable.

```
On Windows:
%APPDATA%\Code\User\globalStorage\sigbots.pros\install\pros-cli-windows
%APPDATA%\Code\User\globalStorage\sigbots.pros\install\pros-toolchain-windows\usr
%USERPROFILE%\cmapi-application

On MacOS:
~/Library/Application\ Support/Code/User/globalStorage/sigbots.pros/install/pros-cli-macos
~/Library/Application\ Support/Code/User/globalStorage/sigbots.pros/install/pros-toolchain-macos
~/cmapi-application

On Linux:
~/.config/Code/User/globalStorage/sigbots.pros/install/pros-cli-linux
~/.config/Code/User/globalStorage/sigbots.pros/install/pros-toolchain-linux/usr
~/cmapi-application
```

#### On Windows

You need to add them to your PATH manually. Please pay a visit to the following website:

https://gist.github.com/nex3/c395b2f8fd4b02068be37c961301caa7

https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/

https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux-unix

#### On MacOS

Run `nano ~/.zshrc` to open the `.zshrc` file. Add the following lines to the end of the file:

```
export PATH="$HOME/Library/Application Support/Code/User/globalStorage/sigbots.pros/install/pros-cli-macos":$PATH
export PROS_TOOLCHAIN="$HOME/Library/Application Support/Code/User/globalStorage/sigbots.pros/install/pros-toolchain-macos"
export PATH="$HOME/cmapi-application:$PATH"
```

Then, save the file with `Ctrl+O`, `enter`, and exit with `Ctrl+X`.

#### On Linux

Run `nano ~/.bashrc` to open the `.bashrc` file. Add the following lines to the end of the file:

```
export PATH="$HOME/.config/Code/User/globalStorage/sigbots.pros/install/pros-cli-linux":$PATH
export PROS_TOOLCHAIN="$HOME/.config/Code/User/globalStorage/sigbots.pros/install/pros-toolchain-linux/usr"
export PATH="$HOME/cmapi-application:$PATH"
```

Then, save the file with `Ctrl+O`, `enter`, and exit with `Ctrl+X`.

Make sure you know what is your default shell of your terminal. You need to modify the `.bashrc` file if you are using bash, or modify the `.zshrc` file if you are using zsh.

### 7. One More Thing on Linux

If you are using Linux, the PROS-CLI installed by the extension might only work with the extension but not on the terminal (SSL issue). Therefore, It is recommended to install the PROS-CLI manually using pip.

```bash
pip3 install pros-cli

# To check if it is installed successfully
which pros
```

On Linux, you also need to add yourself to the `dialout` group before you can upload to the V5 Brain. You then need to log out and log back in for the changes to take effect.

```bash
sudo usermod -a -G dialout $USER
```

## Run From Source

If you want build this tool locally, you can clone the repo and build it with the following commands.

### Build:
```shell
go build -o bin/cmapi-cli.exe
```

### Test coverage:
```shell
go test -coverprofile=coverage
go tool cover -html=coverage
```
