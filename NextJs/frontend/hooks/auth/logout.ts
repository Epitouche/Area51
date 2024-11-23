"use server";
import { cookies } from "next/headers";

if (!process.env.BACKEND_API_URL) {
  throw new Error("BACKEND_API_URL is not defined");
}

export async function logoutCall() {
  const cookie = await cookies();
  const refreshToken = cookie.get('refreshToken')?.value;
  const response = await fetch(`${process.env.BACKEND_API_URL}auth/logout`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ refreshToken }),
  });


  if (!response.ok) {
    return false;
  }

  cookie.delete("token");
  cookie.delete("refreshToken");
  cookie.delete("uuid");
  return true;
}