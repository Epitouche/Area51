"use client";
import { loginCall } from "@hooks";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

interface FormValues {
  email: string;
  password: string;
}

export default function Login() {
  const [error, setError] = useState<string | null>(null);
  const [isClicked, setIsClicked] = useState<boolean>(false);
  const [form, setForm] = useState<FormValues>({
    email: "",
    password: "",
  });

  const router = useRouter();

  const handleSubmit = () => {
    if (!form.email || !form.password) {
      setError("All fields are required");
      return;
    }
    loginCall(form.email, form.password)
      .then((success) => {
        if (!success) {
          setError("Nom d'utilisateur ou mot de passe incorrect");
        } else {
          setError(null);
          router.push("/client");
        }
      })
      .catch(() => {
        setError("An error occurred");
      });
    setForm({ email: "", password: "" });
  };

  useEffect(() => {
    if (form.email != "" && form.password != "") {
      setIsClicked(true);
    }
  }, [form]);

  return (
    <div
      style={{
        display: "flex",
        justifyContent: "center",
        alignItems: "center",
        minHeight: "100vh",
      }}
    >
      <div
        style={{
          boxShadow: "0px 4px 6px rgba(0, 0, 0, 0.1)",
          padding: "2rem",
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
        }}
      >
        <h3 style={{ margin: "0.5rem" }}>Login</h3>

        <input
          type="text"
          placeholder="Email"
          value={form.email}
          onChange={(e) => setForm({ ...form, email: e.currentTarget.value })}
          style={{ marginBottom: "1rem", padding: "0.5rem", width: "100%" }}
        />

        <input
          type="password"
          placeholder="Password"
          value={form.password}
          onChange={(e) =>
            setForm({ ...form, password: e.currentTarget.value })
          }
          style={{ marginBottom: "1rem", padding: "0.5rem", width: "100%" }}
        />

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
          Submit
        </button>

        {error && (
          <div style={{ color: "red", marginTop: "1rem" }}>{error}</div>
        )}
      </div>
    </div>
  );
}
