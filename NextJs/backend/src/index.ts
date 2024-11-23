import express from "express";
import cors from "cors";
import authRouter from "./auth/index";
import "dotenv/config";

const app = express();
const port = 3001;

if (!process.env.BACKEND_API_URL) {
  console.error("Please set the BACKEND_API_URL environment variable");
  process.exit(1);
}

if (!process.env.FRONTEND_API_URL) {
  console.error("Please set the FRONTEND_API_URL environment variable");
  process.exit(1);
}

app.use(
  cors({
    origin: `${process.env.FRONTEND_API_URL}`,
    methods: ["GET", "POST", "PUT", "DELETE"],
    allowedHeaders: ["Content-Type", "Authorization"],
  })
);

app.use(express.json());
app.use("/auth", authRouter);

app.listen(port, () => {
  console.log(`Server is running on ${process.env.BACKEND_API_URL}`);
});
