"use server";
import { cookies } from "next/headers";

if (!process.env.BACKEND_API_URL) {
  throw new Error("BACKEND_API_URL is not defined");
}

export async function loginCall(email: string, password: string) {
  const response = await fetch(`${process.env.BACKEND_API_URL}auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ email, password }),
  });

  if (!response.ok)
    return false;

  return true;
}