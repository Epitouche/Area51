# Mobile

Welcome to the mobile documenation where you will find everything related to the mobile version of the project. \
If you haven't already, check out the POC of mobile and the comparative study documentation [*here*](./POC.md#mobile).

## Figma / Accessibility

If you haven't seen the frontend documentation page yet, go check it out [*here*](./Frontend.md) to see the [*Figma*](./Frontend.md#figma--ui) and the [*Accessibility*](./Frontend.md#accessibility) parts that goes the same for the mobile version of the website.

## React Native

For our mobile development, we chose to use *React Native*. React Native is a popular framework for building cross-platform mobile applications using JavaScript and React. It allows us to write a single codebase that runs on both iOS and Android, significantly reducing development time and effort.

React Native provides a rich set of components and APIs that enable us to create a native-like user experience. Its hot-reloading feature allows for rapid development and iteration, making it easier to test and debug the application. Additionally, React Native has a strong community and a vast ecosystem of libraries and tools, which helps us to quickly find solutions and integrate third-party services.

By using React Native, we can use our existing knowledge of React and JavaScript to build high-quality mobile applications. The framework's performance and flexibility make it an excellent choice for our project, allowing us to deliver a seamless and responsive user experience across multiple platforms.

## Architecture

For the architecture, we decided to adopt a simple yet efficient structure to split and use the mobile version on IOS and android.

The key directories in our architecture are:
- **src**: This is the main directory to code in. You have all the React Native code in this directory:
    - *components*: Where all the components are: ServiceCard, IPInput...
    - *screens*: All the main website pages like Home, Login, Register...
    - *types*: All the types that we use to get info from the backend side are there.
- **ios**: It is where all the files and directories for the IOS to function correclty.
- **android**: Same as the IOS directory. This is where all the important files, for the android to run correclty, are.

To properly use React Native we ensure everything is clean and organize so we can developp our features the best way we can.

### Coding Guidelines

- Follow the existing code style and structure.
- Write clear and concise commit messages.
- Ensure your code is well-documented.
