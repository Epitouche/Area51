"use server";
import { cookies } from "next/headers";

export async function useGetCookies(type: string) {
  const cookie = await cookies();
  if (type === "uuid")
    return cookie.get("uuid");
  else if (type === "token")
    return cookie.get("token");
  else if (type === "refreshToken")
    return cookie.get("refreshToken");

  return cookie.getAll();
}
