# pom

Maven POM parser written in Go that supports recursive variable expansion/interpretation.

## Library

`pom.xml`:
```xml
<project>
	<version>${p1}.${a.b}.z</version>
	<properties>
		<p1>x${p.b}</p1>
		<p.b>x</p.b>
	</properties>
</project>
```

```go
model, err := Unmarshal([]byte(data))

// supply additional variables
model.SetProperty("a.b", "y")

version, err := model.Get("version")

// version is xx.y.z
```

## Program

`pom.xml`:
```xml
<?xml version="1.0" encoding="UTF-8" ?>
<project
        xmlns="http://maven.apache.org/POM/4.0.0"
        xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
        xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/maven-v4_0_0.xsd">

    <version>${version.major}.${version.minor}.${version.patch}-bitbucket${bitbucket.major}</version>

    <properties>
        <version.major>0</version.major>
        <version.minor>1</version.minor>
        <version.patch>5</version.patch>

        <bitbucket.major>6</bitbucket.major>
        <bitbucket.minor>2</bitbucket.minor>
        <bitbucket.patch>0</bitbucket.patch>
        <bitbucket.version.generated>${bitbucket.major}.${bitbucket.minor}.${bitbucket.patch}</bitbucket.version.generated>
    </properties>
</project>
```

Run:
```sh
$ pom version < pom.xml
```

Result:
`0.1.5-bitbucket6`

## Installation

```
go get github.com/reconquest/pom/cmd/pom
```

## LICENSE

MIT
