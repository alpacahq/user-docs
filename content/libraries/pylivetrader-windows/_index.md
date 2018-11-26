---
title: Setup pylivetrader on Windows
weight: 50
---

# Setup pylivetrader on Windows

*These steps are from a Windows user who kept track while successfully setting up pylivetrader to run a Quantopian-like algorithm on Alpaca*

## OS
This was on Windows 7 and should also apply to Windows 10.

## Visual Studio and .NET
Prerequisites are Python 3.6 or higher and some form of C++ compiler, which requires .NET Framework.

**C++ compiler install**
Various options are available to have a compiler available, I installed Visual Studio Community 2017 Preview (free).
VS Code at the end of Anaconda might also work, not sure (.NET Framework is still needed for that).

If a C++ compiler were not present, this is the error that would occur on `pip install pylivetrader`

```
Running setup.py clean for lru-dict
  error: Microsoft Visual C++ 14.0 is required.
Get it with "Microsoft Visual C++ Build Tools":
http://landinghub.visualstudio.com/visual-cpp-build-tools, however instead of Build Tools, I'm opting for Visual Studio Community below.
```

(Note that link cited in the error message is 404, try instead https://visualstudio.microsoft.com/visual-cpp-build-tools/)

Installed **.NET Framework** (needed for Visual Studio Community)

- https://www.microsoft.com/net/download/dotnet-framework-runtime goes to (direct link ...)
[https://www.microsoft.com/net/download/thank-you/net472](https://www.microsoft.com/net/download/thank-you/net472 ".NET Framework")

Installed **Visual Studio Community**
https://visualstudio.microsoft.com/downloads/ goes to (direct link ...)
[https://visualstudio.microsoft.com/thank-you-downloading-visual-studio/?ch=pre&sku=Community&rel=15](https://visualstudio.microsoft.com/thank-you-downloading-visual-studio/?ch=pre&sku=Community&rel=15 "Visual Studio Community")

During the install, I selected these options:

- Desktop development with C++
- Python development

## Install Python 3.6

  One option is python.org: https://www.python.org/downloads/release/python-370/

  I chose Anaconda (Python package) here:

  [Anaconda 64-Bit Graphical Installer](https://repo.anaconda.com/archive/Anaconda3-5.2.0-Windows-x86_64.exe "Anaconda") (631 MB) ... installing that to `C:\Python36`

Corrected **Path**

In the Windows path, my previous `C:\python27` had not been removed. To correct the path (a list of locations Windows looks to, for finding certain things), I did the following:

Start button > Computer > right click > Properties > Advanced system settings > Environment Variables > *User* > Path
... setting that to:
C:\Python36;C:\Python36\Library\mingw-w64\bin;C:\Python36\Library\usr\bin;C:\Python36\Library\bin;C:\Python36\Scripts;C:\Program Files\Docker Toolbox

Start button > Computer > right click > Properties > Advanced system settings > Environment Variables > *System* > Path
%SystemRoot%\system32;%SystemRoot%;%SystemRoot%\System32\Wbem;C:\Program Files\Git\cmd

Not certain about that but eventually typing `path` in the command prompt, the path wound up like this:

```
PATH=C:\Python36;C:\Python36\Library\mingw-w64\bin;C:\Python36\Library\usr\bin;C:\Python36\Library\bin;C:\Python36\Scripts;C:\Windows\system32;C:\Windows;C:\Windows\
System32\Wbem;C:\Program Files\Git\cmd;C:\Program Files\Microsoft VS Code\bin;C:\Program Files\Docker Toolbox
```

*To open a command prompt (command window):
Start > Run > cmd
or Windows key > r* > cmd

## Install Python modules and pylivetrader

Updated **pip** (python tool for installing things)
`python -m pip install --upgrade pip`

Installed **pylivetrader** ([link](https://github.com/alpacahq/pylivetrader "pylivetrader"))

`pip install pathlib`

`pip install pylivetrader`

Message says:
`Successfully built lru-dict`
(because the C++ compiler ok)

Installed **pipeline-live** ([link](https://github.com/alpacahq/pipeline-live "pipeline-live"))

`pip install pipeline-live`

## Set up Keys and Configure
Obtain key from [Dashboard](https://app.alpaca.markets/dashboard/overview)

The slider at the top of the page switches between paper trading and live, and results in different keys for the two separate purposes.

The command to start trading includes an option/switch like `--backend-config ap.yaml` to specify the keys that will be used for live or paper. I chose the live trading yaml file name `ap.yaml`. My paper trading yaml file is called `paper.yaml`.

Contents of the yaml files below.
*These keys are trimmed for the example*
Comments in the file are the `set` commands to also set the keys as environment variables. Notice on paper trading, the env variable `APCA_API_BASE_URL` is different from the yaml file `base_url` currently.

**Keys Live**

```
key_id: AKxxxxxx
secret: FnvMxxxxxx

 # Live
 # set APCA_API_KEY_ID=AKxxxxxx
 # set APCA_API_SECRET_KEY=FnvMxxxxxx
```

**Keys Paper Trading**

```
key_id: KMxxxxxx
secret: Fvtxxxxxx
base_url: https://paper-api.alpaca.markets

 # Paper
 # set APCA_API_BASE_URL=https://paper-api.alpaca.markets
 # set APCA_API_KEY_ID=KMxxxxxx
 # set APCA_API_SECRET_KEY=Fvtxxxxxx
```

Make an **algorithm**. Find an example online.

## Start algorithm at any time of day

```
C:\ap> pylivetrader.exe run -f alpaca_test1.py --backend-config paper.yaml
[2018-09-16 21:31:33.615813] INFO: Algorithm: livetrader start running with backend = alpaca data-frequency = minute
```