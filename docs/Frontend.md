# Frontend

Welcome to the frontend documentation where you will find everything related to frontend side of the project.

If you haven't already, check out the POC of frontend and the comparative study documentation [*here*](./POC.md#frontend).

## Figma / UI

Right at the beginning we first started by making a **figma** to center all of our ideas in one place to have a good looking UI for our website. We spent around a week and a half on the figma to really know what our website was going to be like and to show everything for our **Planning** first notation of the project (Defense as we call it).

Our figma is split into three distinct but connected part:
- **Design** : Where all the main designs are like buttons or even graphical charts
- **Web** : All the websites pages and components like the Navbar or the login page
- **mobile** : Same as the website part but for mobile.

Check out our [*Figma*](https://www.figma.com/design/SDi5Wr1talXN5o4wirUuSD/AREA-UI%2FUX?node-id=8-12&p=f&t=CF4lYEruvTCz68Un-0) to see it by yourself.

## Nuxt.js / Tailwind CSS

As you may have already seen in the POC of frontend, we chose to use Nuxt.js in combination with Tailwind CSS. Nuxt.js, is a powerful framework built on top of Vue.js, provides us with a robust structure for building server-side rendered applications and static sites. Its features, such as automatic routing, server-side rendering, and static site generation, align perfectly with our project requirements.

Tailwind CSS, a utility-first CSS framework, allows us to rapidly build custom user interfaces without writing a lot of custom CSS. By using Tailwind CSS, we can maintain a consistent design system across our application and easily adapt to design changes. The combination of Nuxt.js and Tailwind CSS enables us to build a performant, scalable, and maintainable frontend.

## Architecture

For the architecture, we decided to adopt a simple yet efficient structure provided by Nuxt.js. This includes the default directory structure, which helps in organizing our code and maintaining a clean project layout.

The key directories in our architecture are:
- **pages**: Contains all the page components, which Nuxt.js automatically sets up with routing.
- **components**: Houses reusable Vue components used across different pages, such as Navbar, Button, Input...
- **assets**: Stores static assets like images and stylesheets.

By using Nuxt.js's conventions and features, we ensure that our project remains organized and scalable, with minimal configuration required.

## Accessibility

Ensuring our application is accessible to all users, including those with disabilities, is a top priority. We have implemented several measures to enhance accessibility for users with various needs, including those who are hard of hearing, visually impaired, or have motor disabilities.

### Visual Accessibility

For users with visual impairments, we have taken the following steps:
- **Color Contrast**: We ensure that text and background colors have sufficient contrast to be easily readable. This helps users with low vision or color blindness.
- **Text Size and Scaling**: Our application supports text resizing and scaling, allowing users to adjust text size according to their preferences.
- **ARIA Roles and Attributes**: We use ARIA (Accessible Rich Internet Applications) roles and attributes to provide additional context to screen readers, making it easier for visually impaired users to navigate the application.

### Hearing Accessibility

For users who are hard of hearing or deaf, we have implemented the following measure:
- **Visual Indicators**: We use visual indicators, such as icons and text, to convey important information that might otherwise be communicated through sound.

### Motor Accessibility

For users with motor disabilities, we have taken the following steps:
- **Keyboard Navigation**: Our application is fully navigable using a keyboard. We ensure that all interactive elements, such as buttons and links, can be accessed and operated using keyboard shortcuts.
- **Focus Management**: We manage focus states effectively, ensuring that users can easily see which element is currently focused and navigate through the application without losing context.
- **Accessible Forms**: All form elements are properly labeled, and we provide clear instructions and error messages to assist users in completing forms accurately.

### Cognitive Accessibility

For users with cognitive disabilities, we have implemented the following measures:
- **Simple Language**: We use clear and simple language throughout the application to make it easier for users to understand the content.
- **Consistent Navigation**: Our navigation structure is consistent and predictable, helping users to easily find their way around the application.
- **Error Prevention and Recovery**: We provide helpful error messages and guidance to assist users in recovering from mistakes and completing tasks successfully.

By incorporating these accessibility features, we aim to create an inclusive and user-friendly application that can be used by everyone, regardless of their abilities.
