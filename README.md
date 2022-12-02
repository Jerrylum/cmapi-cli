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

### 5. Add PROS-CLI to PATH

The PROS-CLI should be installed at the following location by the extension:

```
On Windows:
%APPDATA%\Code\User\globalStorage\sigbots.pros\install\pros-cli-windows

On Linux:
~/.config/Code/User/globalStorage/sigbots.pros/install/pros-cli-linux
```

You need to add it to your PATH manually.

On Windows: https://www.architectryan.com/2018/03/17/add-to-the-path-on-windows-10/

On Linux: https://stackoverflow.com/questions/14637979/how-to-permanently-set-path-on-linux-unix

However, if you are using Linux, the PROS-CLI installed by the extension might only work with the extension but not on the terminal (SSL issue). Therefore, It is recommended to install the PROS-CLI manually using pip.

```bash
pip3 install pros-cli

which pros
```

### 6. Set PROS_TOOLCHAIN Environment Variable

The toolchain should be installed at the following location by the extension:

```
On Windows:
%APPDATA%\Code\User\globalStorage\sigbots.pros\install\pros-toolchain-windows\usr

On Linux:
~/.config/Code/User/globalStorage/sigbots.pros/install/pros-toolchain-linux/usr
```

You need to set PROS_TOOLCHAIN Environment Variable to the location of the toolchain. It is similar to adding the CLI to the PATH environment variable.

### 7. Install CMAPI-CLI

First, download the most recent build for your OS from the releases page. Then, rename the file to `cmapi-cli`. Now, put the executable in a folder that is in your PATH environment variable or add the folder to your PATH environment variable.

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
