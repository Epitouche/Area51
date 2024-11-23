"use client";
import "dotenv/config";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function Home() {
  const [isClicked, setIsClicked] = useState<boolean>(true);
  const router = useRouter();

  const handleSubmit = () => {
    router.push("/login");
  };


  return (
    <div style={{ alignContent: "center", justifyItems: "center" }}>
      <div style={{ paddingTop: "3rem", alignContent: "center" }}>
        <text style={{ fontSize: "8rem" }}>HOME</text>
        <button
          type="submit"
          onClick={() => handleSubmit()}
          disabled={!isClicked}
          style={{
            padding: "0.75rem",
            width: "100%",
            backgroundColor: isClicked ? "#4CAF50" : "#ccc",
            color: "white",
            border: "none",
            borderRadius: "4px",
          }}
        >
          login
        </button>
      </div>
    </div>
  );
}
