# Developer Documentation for Frontend Project

## Introduction
This document provides a detailed overview of the frontend project, built using **Nuxt.js**. It serves as a guide for developers to understand, maintain, and contribute to the project effectively.

### Key Features
- **Nuxt.js Framework**: A Vue.js framework for building server-side rendered (SSR) applications.
- **TypeScript Support**: Strong typing for better maintainability.
- **Pinia**: For state management.
- **ESLint & Prettier**: For code linting and formatting.
- **Devcontainer Support**: Development environment configured with a devcontainer for consistency and ease of setup.

---

## Prerequisites

To work on this project, ensure you have the following installed:

- **Docker**: For running the devcontainer ([Download Docker](https://www.docker.com/)).
- **VS Code**: Recommended IDE with the Remote - Containers extension installed ([Install Extension](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers)).

---

## Setting Up the Project

### Using the Devcontainer

1. **Clone the Repository:**
   ```bash
   git clone <repository-url>
   cd frontend
   ```

2. **Open the Project in VS Code:**
   ```bash
   code .
   ```

3. **Reopen in Devcontainer:**
   - Press `CTRL` + `SHIFT` + `P` in VS Code.
   - Search for and select: `Dev Containers: Rebuild and Reopen in Container`.

5. **Access the Application:**
   Open [http://localhost:8081](http://localhost:8081) in your browser.

---

## Project Structure

The project is organized as follows:

```
frontend/
├── assets/         # Static assets (images, styles)
├── components/     # Reusable Vue components
├── layouts/        # Application layouts
├── pages/          # Page components (routes are auto-generated)
├── plugins/        # Plugins to initialize before app load
├── public/         # Static files served at the root URL
├── stores/         # State management (Pinia)
├── app.vue         # Main app component
├── nuxt.config.js  # Nuxt configuration
├── package.json    # Project dependencies and scripts
├── tsconfig.json   # TypeScript configuration
└── README.md       # Project overview
```

---

## Scripts

The following npm scripts are available:

- **`npm run dev`**: Starts the development server.
- **`npm run build`**: Builds the application for production.
- **`npm run start`**: Runs the production server.
- **`npm run lint`**: Lints the code using ESLint.
- **`npm run lint:fix`**: Fixes linting errors automatically.

---

## Configuration

### Environment Variables

Environment variables are stored in a `.env` file. Here are some key variables:

- `API_BASE_URL`: Base URL for backend API.
- `NODE_ENV`: Application environment (development, production).

To create a `.env` file, use the provided `.env.example` as a template:
```bash
cp .env.example .env
```

### ESLint & Prettier

The project enforces coding standards using **ESLint** and **Prettier**. Ensure your code passes linting before committing changes:
```bash
npm run lint
```

---

## Contribution Guidelines

1. **Branching Strategy:**
   - Use the `frontend` branch for base code in your feature branches.
   - Create feature branches for new features or bug fixes.
     ```
     git checkout -b feature/your-feature-name
     ```
   - Link the issue number in the branch name for reference. 

2. **Commit Messages:**
   Follow the [Conventional Commits](https://docs.google.com/document/d/1JRWCsIwZGD9q2ZuTiY117qcswexsDutevPy90z75hJ8/edit?usp=sharing) format:

3. **Pull Requests:**
   Follow the [Convensional Commits](https://docs.google.com/document/d/1B3cneL52bKgU3n2Bmr2hV2LLxdBvaoqEs8ddio9EDUY/edit?usp=sharing) format:
   - Submit pull requests for review before merging.
   - Ensure the CI/CD pipeline passes.

---

## Troubleshooting

### Common Issues

1. **Dependencies not installed:**
   Ensure you have run:
   ```bash
   npm install
   ```

2. **Port already in use:**
   Stop any other service running on port 8081 or change the port in `nuxt.config.js`.

3. **ESLint errors:**
   Run:
   ```bash
   npm run lint:fix
   ```

---

## Useful Links

- [Nuxt.js Documentation](https://nuxtjs.org/docs)
- [Vue.js Documentation](https://vuejs.org/)
- [Pinia Documentation](https://pinia.vuejs.org/)
- [Devcontainer Documentation](https://code.visualstudio.com/docs/remote/containers)
