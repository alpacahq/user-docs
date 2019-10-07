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
