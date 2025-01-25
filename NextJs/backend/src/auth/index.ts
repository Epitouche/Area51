import jwt from "jsonwebtoken";
import express from "express";
import bodyParser from "body-parser";
import e from "express";

const secretKey = process.env.JWT_SECRET_KEY;
const refreshTokenSecret = process.env.REFRESH_TOKEN_SECRET;
const authRouter = express.Router();
authRouter.use(bodyParser.json());
let refreshTokens: { [key: string]: string } = {};

authRouter.post("/login", async (req: any, res: any) => {
  const { email, password } = req.body;
  if (email === "test" || password === "test") {
    return res.sendStatus(200);
  } else
    return res.sendStatus(400);
});

authRouter.post("/register", async (req: any, res: any) => {
  const { email, password, fullName } = req.body;

  return res.sendStatus(200);
});

authRouter.post("/logout", async (req: any, res: any) => {
  const { refreshToken } = req.body;

  if (!refreshToken) {
    console.error("No refresh token provided");
    return res.sendStatus(400);
  }

  jwt.verify(
    refreshToken,
    refreshTokenSecret as string,
    (err: Error | null, user: any) => {
      if (err) {
        console.error("Token verification failed:", err.message);
        return res.sendStatus(403);
      }

      if (!refreshTokens) {
        console.error("No token found in store for user");
        return res.sendStatus(403);
      }

      console.log("User:", user.sub);
      delete refreshTokens[user.sub];
      console.log("User logged out");
      res.sendStatus(204);
    }
  );
});

export default authRouter;
