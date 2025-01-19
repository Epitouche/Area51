# How to Sign Your React Native App

## Prerequisites

Before getting started, ensure you have the following installed on your system:

- **Java Development Kit (JDK)**: Java Development Kit (version 11 or newer). [Download here](https://www.oracle.com/java/technologies/javase-jdk11-downloads.html)

## Generating a Signing Key

### Step 1: Create a Key Store

1. Open a terminal and run the following command to generate a key store:
    ```sh
    keytool -genkeypair -v -storetype PKCS12 -keyalg RSA -keysize 2048 -validity 10000 -keystore my-release-key.keystore -alias my-key-alias
    ```

2. You will be prompted to enter some information about the key. Fill out the necessary details:
    - **Keystore password**: A secure password for the keystore.
    - **Re-enter new password**: Confirm the password.
    - **What is your first and last name?**
    - **What is the name of your organizational unit?**
    - **What is the name of your organization?**
    - **What is the name of your City or Locality?**
    - **What is the name of your State or Province?**
    - **What is the two-letter country code for this unit?**
    - **Is CN=..., OU=..., O=..., L=..., ST=..., C=... correct?**: Answer `yes`.

This will create a file called `my-release-key.keystore` in your current directory.

### Step 2: Setting Up Gradle Variables

1. Place the `my-release-key.keystore` file under the `android/app` directory of your React Native project.

2. Edit the `~/.gradle/gradle.properties` file (create this file if it doesn't exist) and add the following lines to it:
    ```properties
    MYAPP_UPLOAD_STORE_FILE=my-release-key.keystore
    MYAPP_UPLOAD_KEY_ALIAS=my-key-alias
    MYAPP_UPLOAD_STORE_PASSWORD=your-keystore-password
    MYAPP_UPLOAD_KEY_PASSWORD=your-key-password
    ```

### Step 3: Configuring Gradle for Signing

1. Edit the `android/app/build.gradle` file in your React Native project to add signing configuration:
    ```gradle
    android {
        ...
        defaultConfig { ... }
        signingConfigs {
            release {
                storeFile file(MYAPP_UPLOAD_STORE_FILE)
                storePassword MYAPP_UPLOAD_STORE_PASSWORD
                keyAlias MYAPP_UPLOAD_KEY_ALIAS
                keyPassword MYAPP_UPLOAD_KEY_PASSWORD
            }
        }
        buildTypes {
            release {
                ...
                signingConfig signingConfigs.release
            }
        }
    }
    ```

## Building the Signed APK

1. Navigate to the `android` directory of your React Native project and run:
    ```sh
    cd android
    ./gradlew assembleRelease
    ```

2. After the build completes, you can find the signed APK at `android/app/build/outputs/apk/release/app-release.apk`.

## Conclusion

You have successfully signed your React Native application. You can now distribute your signed APK to users or upload it to the Google Play Store.