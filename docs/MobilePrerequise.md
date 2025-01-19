# How to Start the React Native App (without Expo)

## Prerequisites

Before getting started, ensure you have the following installed on your system:

- **Node.js**: [Download and install Node.js](https://nodejs.org/)
- **Watchman**: [Installation guide for Watchman](https://facebook.github.io/watchman/docs/install.html) (only for macOS)
- **JDK**: Java Development Kit (version 11 or newer). [Download here](https://www.oracle.com/java/technologies/javase-jdk11-downloads.html)
- **Android Studio**: [Download Android Studio](https://developer.android.com/studio)

## Setting Up the Environment

### Step 1: Install Node.js and Watchman

Install Node.js from the official website. Watchman is required only for macOS and can be installed via Homebrew:

```sh
brew install watchman
```

### Step 2: Install the JDK

Download and install the Java Development Kit (JDK) from Oracle's website. Make sure to set the `JAVA_HOME` environment variable.

### Step 3: Install Android Studio

1. **Download Android Studio**:
    - Go to the [Android Studio download page](https://developer.android.com/studio) and install it.

2. **Install Android SDK**:
    - Open Android Studio.
    - Go to `Preferences` > `Appearance & Behavior` > `System Settings` > `Android SDK`.
    - Select `SDK Tools` tab and ensure the following items are checked:
      - Android SDK Build-Tools
      - Android SDK Platform-Tools
      - Android Emulator
      - Android SDK Tools

3. **Set up environment variables**:
    - Add the following lines to your `.bash_profile`, `.zshrc`, or equivalent:

```sh
export ANDROID_HOME=$HOME/Library/Android/sdk
export PATH=$PATH:$ANDROID_HOME/emulator
export PATH=$PATH:$ANDROID_HOME/tools
export PATH=$PATH:$ANDROID_HOME/tools/bin
export PATH=$PATH:$ANDROID_HOME/platform-tools
```

4. **Reload your terminal**:

```sh
source ~/.bash_profile
```

## Creating a React Native Project

1. **Initialize a new React Native project**:

```sh
npx react-native init MyApp
cd MyApp
```

2. **Run the app on Android**:

You have two options here:

### Option 1: Using an Android Emulator
- Make sure you have an Android emulator running.
- Then run:
  ```sh
  npx react-native run-android
  ```

### Option 2: Using a Physical Android Device
- **Enable USB Debugging on your device**:
  - Go to `Settings` > `About phone` and tap `Build number` seven times to enable developer options.
  - Go to `Settings` > `Developer options` and enable `USB debugging`.

- **Connect your device to your computer**:
  - Use a USB cable to connect your Android device to your computer.
  - Verify the connection by running:
  ```sh
  adb devices
  ```
  - Your device should be listed.

- **Run the app on the device**:
  ```sh
  npx react-native run-android
  ```

## Run the app on iOS (macOS only):

Make sure you have Xcode installed and then run:

```sh
npx react-native run-ios
```