"use server";

import { cookies } from 'next/headers'

if (!process.env.BACKEND_API_URL) {
  throw new Error("BACKEND_API_URL is not defined");
}

export async function registerCall(
  password: string,
  company: string,
  fullName: string,
  serial_number: string,
  role: string
) {
  const cookie = await cookies();
  const response = await fetch(`${process.env.BACKEND_API_URL}auth/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({ serial_number, password, role, company, fullName }),
  });

  if (!response.ok) {
    return false;
  }

  const data = await response.json();
  console.log("Data:", data);
  if (data.error) return false;

  cookie.set("token", data.token);
  cookie.set("refreshToken", data.refreshToken);
  cookie.set("uuid", data.uuid);
  return true;
}
