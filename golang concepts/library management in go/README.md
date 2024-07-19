Library management in go revolves around repository, modules and packages.

More then one modules in one repository is not advisable.

A module can contain multiple packages inside it.

Before doing anything we have to declare that this is our module. every module has Globally unique identifier. By using it, it can be downloaded.


Summary :
1. Go code is grouped into packages and packages are grouped into modules.

2. A module specifies the dependencies needed to run our code, including the go version and set of other modules it requires in the go.mod file.

