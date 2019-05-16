---
title: Install Pylivetrader on Windows
weight: 50
---

Two ways to install pylivetrader on Windows.  
The first section uses Anaconda environments.  
Or jump to ![](#without-conda-environments "Setup without using Anaconda environments")

------------------------------------------------
# Install Pylivetrader on Windows
## _Using Anaconda Environment_
------------------------------------------------

Pylivetrader packages for Windows are hosted on our channel on Anaconda. To get started, follow these steps.

1. Install the Python 3.7 version of [Anaconda](https://www.anaconda.com/download/#windows) if you don’t already have it. You may be prompted to register an Anaconda account after installation.
2. From your start menu, open Anaconda Prompt.
3. Navigate into a development folder.

    ```
    mkdir alpacahq
    cd alpacahq
    ```

4. Add the conda-forge, Quantopian, and alpacahq channels.

    ```
    conda config --add channels conda-forge
    conda config --add channels Quantopian
    conda config --add channels alpacahq
    ```

5. Create a virtual environment with the necessary packages installed.

    ```
    conda create -n pylivetraderenv pylivetrader
    ```

6. Activate your new virtual environment.

    ```
    source activate pylivetraderenv
    ```

Once you follow these steps, you should be able to run pylivetrader from the Anaconda Prompt terminal.
```
pylivetrader run exampleAlgorithm.py --backend-config exampleConfig.yaml
```
You can reactivate this environment at any time by running `source activate pylivetraderenv` inside the Anaconda prompt.

If you encounter issues with setting your environment up, feel free to ask for help in our community Slack. A link to join is in the sidebar.

------------------------------------------------
# Setup of pylivetrader on Windows  
## _without conda environments_
------------------------------------------------
Anaconda is used here to place files but then `conda` is not used, just standard Python like `pip`.

*These steps are from a Windows user who kept track while successfully setting up pylivetrader to run a Quantopian-like algorithm on Alpaca.*

### Python, using Anaconda setup  
I recommend Python 3.6 for now, to avoid potential problems. This is it, in Anaconda:  
`Anaconda3-5.2.0-Windows-x86_64.exe 	631.3M 	2018-05-30 13:04:18 	62244c0382b8142743622fdc3526eda7`  
https://repo.continuum.io/archive/Anaconda3-5.2.0-Windows-x86_64.exe  
&nbsp; &nbsp;    from https://repo.continuum.io/archive/  
&nbsp; &nbsp;    Install as Just Me or All Users  
&nbsp; &nbsp;      to `C:\Python36` _(short simple paths make easier reading, copy/paste and everything else along the way)_  
    
&nbsp; &nbsp;    Check the boxes: Add to path and Register as default 
    
&nbsp; &nbsp;    Install VSCode at the end too, it can be useful (although it won't break in to breakpoints on dependent files).

Open a new `cmd` prompt and verify path. I added some newlines here for clarity, showing my path as an example:
```
C:\> path
    PATH=C:\windows\system32;C:\windows;C:\windows\System32\Wbem;C:\windows\System32\WindowsPowerShell\v1.0\;C:\windows\System32\OpenSSH\;
    C:\Program Files\Microsoft VS Code\bin;
    C:\Python36;
    C:\Python36\Library\mingw-w64\bin;
    C:\Python36\Library\usr\bin;
    C:\Python36\Library\bin;
    C:\Python36\Scripts;
    C:\Users\<user>\AppData\Local\Microsoft\WindowsApps	
```

### Visual C++ 2015 Build Tools
**This is a key step that resolves the `lru-dict` compile snag.**  
I found the answer here:  
&nbsp; &nbsp;    https://stackoverflow.com/questions/44951456/pip-error-microsoft-visual-c-14-0-is-required  
&nbsp; &nbsp;    ... in this link (in an answer) and it works. The link is unusual, and includes the period at the end:  
&nbsp; &nbsp; &nbsp; &nbsp;        http://go.microsoft.com/fwlink/?LinkId=691126&fixForIE=.exe.
    
### Keys: Use keyz.cmd (see the contents of the `cmd` file below)  
Obtain keys from https://app.alpaca.markets/paper/dashboard/overview (paper)  
Paste your key and secret key into the file below.  
Run that at `cmd` prompt.  
Keys are now set for the current `cmd` window and all others going forward,  
&nbsp; and any processes that rely on them.

### Pylivetrader  
On the Start menu, open 'Visual C++ 2015 MSBuild Command Prompt' _so the compiler will be found._

Run these:
```
    python.exe -m pip install --upgrade pip
    pip install zipline
    pip install pylivetrader
    pip install pipeline-live
    pip install pipdeptree
    pipdeptree -p pylivetrader
```
    
In that last command, there are likely to be some conflicts, depending on one's own environment.  
For example, mine:
```
    C:\> pipdeptree -p pylivetrader
    Warning!!! Possibly conflicting dependencies found:
    * intervaltree==3.0.2
     - sortedcontainers [required: >=2.0,<3.0, installed: 1.5.10]
    * distributed==1.21.8
     - msgpack [required: Any, installed: ?]
```
     
Those are easy to resolve. I did so as follows. Just use the double-quote pattern for versions:
```
    pip install msgpack
    pip install "sortedcontainers>=2.0"
```

Try also:
```
    pipdeptree -p zipline
```
        
#### Now all set. Everything is clean.

### Run an example to make sure it works:
Grab the example at https://github.com/alpacahq/pylivetrader/blob/master/examples/q01/algo.py  
&nbsp; &nbsp;    ... using the `Raw` button.

Save that as `algo1.py`.  
Open a new `cmd` window, command prompt (console window).  
`cd` to the directory where you saved that `algo1.py` file. I use C:\ap  
[The stock market doesn't need to be open to test this]

Type at the prompt:  
```
pylivetrader run algo1.py
```

This output indicates it is fine:
```
[2019-05-08 01:16:42.080187] INFO: Algorithm: livetrader start running with backend = alpaca data-frequency = minute
93
[2019-05-08 01:16:52.327810] INFO: algo:
Algorithm initialized variables:
 context.MaxCandidates 100
 LowVar 6
 HighVar 40    
```

_That was a starter algorithm by Alapaca, you can ask around in 
&nbsp;  https://alpaca-community.slack.com/ and might find alternatives._
    
### Other
These are some other programs I use *a lot* on Windows for development:

##### Notepad++  
https://notepad-plus-plus.org/  
Make sure the 32-bit version even on x64, for TextFX and Python support.  
There's also a portable version out there if you search the web for it.

TextFX Characters  
&nbsp; &nbsp;    Plugins Admin > 'textfx' > Install  
Settings > Preferences > Language > Python > Tab characters > Replace by 4 spaces

##### CompareIt  
&nbsp; &nbsp;    https://www.grigsoft.com/wincmp3.htm

##### Search & Replace  
&nbsp; &nbsp;    https://www.funduc.com/search_replace.htm
    
##### Paintshop Pro 4.14  
(contact me, it is old. just the essentials)  
For saving screen captures.  
On start, one time:  
&nbsp; &nbsp;    The first time, run it as administrator, right click the icon and dig for that option.  
&nbsp; &nbsp;      (otherwise will say it can't update the registry [with its settings])  
&nbsp; &nbsp;    Clear initial dialogs.  
&nbsp; &nbsp;    Move toolbars all to the same line by dragging just off to the side of buttons, a tiny area available for that.  

##### PyCharm
https://www.jetbrains.com/pycharm/download/#section=windows  
Another program I've been tinkering with, and would be very good for development purposes, except for one thing.  
This can break into breakpoints set in dependent files too (unlike VS Code), not just the algo.  
  However there's no known way for it to initiate debugging with pylivetrader.exe in use, currently.

#### keyz.cmd file content
```
:: Sets Alpaca keys, the environment variables, 
::   for both the current `cmd` window and/or other applications going forward.
:: To obtain keys, visit https://app.alpaca.markets/paper/dashboard/overview
:: I call this file keyz.cmd to be able to just type `keyz` at the prompt after changes
::   because the word `keys` in Windows is already taken.

@echo off

:: Set keys for the current window
set APCA_API_KEY_ID=PK85___YOUR_KEY___6PVN
set APCA_API_SECRET_KEY=6waXwdMaTX___YOUR_KEY___nc1UQqkNlC
set APCA_API_BASE_URL=https://paper-api.alpaca.markets

:: Set them also for other windows and processes going forward
setx APCA_API_KEY_ID     %APCA_API_KEY_ID%
setx APCA_API_SECRET_KEY %APCA_API_SECRET_KEY%
setx APCA_API_BASE_URL   %APCA_API_BASE_URL%

:: Displaying what was just set.
set apca

:: Or for copy/paste manually ...
:: setx APCA_API_KEY_ID     'PK85___YOUR_KEY___6PVN'
:: setx APCA_API_SECRET_KEY '6waXwdMaTX___YOUR_KEY___nc1UQqkNlC'
:: setx APCA_API_BASE_URL   'https://paper-api.alpaca.markets'
```


#### ap_files.cmd (Another script I use and recommend to make things easier)
```
:: I run a little `ap_files.cmd` file like this to copy some of the latest files used by pylivetrader to one place.
:: Then with https://www.funduc.com/ftp/setupsr64.exe (Search & Replace) 
::   I can quickly & easily look thru them all for whatever reason.
:: When the need arises to find something, it saves a ton of time.
:: Search & Replace is one of my main tools used most often, digging thru algorithms for bits code.
:: mkdir will fail if already exists causing no problem.

cd %USERPROFILE%\documents
mkdir ap\site-packages
cd %USERPROFILE%\documents\ap\site-packages

xcopy /ifdryshck c:\python36\Lib\site-packages\alpaca_trade_api     alpaca_trade_api
xcopy /ifdryshck c:\python36\Lib\site-packages\iexfinance           iexfinance
xcopy /ifdryshck c:\python36\Lib\site-packages\pipeline_live        pipeline_live
xcopy /ifdryshck c:\python36\Lib\site-packages\pylivetrader         pylivetrader
xcopy /ifdryshck c:\python36\Lib\site-packages\requests             requests
xcopy /ifdryshck c:\python36\Lib\site-packages\trading_calendars    trading_calendars
xcopy /ifdryshck c:\python36\Lib\site-packages\urllib3              urllib3
xcopy /ifdryshck c:\python36\Lib\site-packages\logbook              logbook
xcopy /ifdryshck c:\python36\Lib\site-packages\zipline              zipline

:: Another step, once those are copied, I delete all of the .pyc files.
:: They are pre-compiled cache versions of .py files for faster startup, not needed there.
del /s /q *.pyc

:: Also deleting test files & directories.
:: Files
del /s /q test*
:: Directories. (Sad that this line needed to be so complex)
for /f %i in ('dir /a:d /s /b *test*') do rd /s /q %i
```

_Section written by G_Ha -- on https://alpaca-community.slack.com/_
