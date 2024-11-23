"use client";
import { registerCall } from "@hooks";
import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";

interface FormValues {
  password: string;
  company: string;
  name: string;
  surname: string;
  serial_number: string;
  role: string;
}

export default function Register() {
  const [error, setError] = useState<string | null>(null);
  const [isClicked, setIsClicked] = useState<boolean>(false);
  const [form, setForm] = useState<FormValues>({
    password: "",
    company: "",
    name: "",
    surname: "",
    serial_number: "",
    role: "",
  });

  const router = useRouter();

  const handleSubmit = () => {
    if (
      !form.password ||
      !form.company ||
      !form.name ||
      !form.surname ||
      !form.serial_number ||
      !form.role
    ) {
      setError("All fields are required");
      return;
    }
    registerCall(
      form.password,
      form.company,
      `${form.name} ${form.surname}`,
      form.serial_number,
      form.role
    )
      .then((success) => {
        if (!success) {
          setError("User already exists");
        } else {
          setError(null);
          router.push("/dashboard");
        }
      })
      .catch(() => {
        setError("An error occurred");
      });
  };

  useEffect(() => {
    if (
      form.company != "" &&
      form.name != "" &&
      form.surname != "" &&
      form.serial_number != "" &&
      form.role != "" &&
      form.password != ""
    ) {
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
        <h3 style={{ margin: "0.5rem" }}>Register</h3>

        <div
          style={{
            display: "flex",
            justifyContent: "space-between",
            width: "100%",
          }}
        >
          <input
            type="text"
            placeholder="Name"
            value={form.name}
            onChange={(e) => setForm({ ...form, name: e.currentTarget.value })}
            style={{ marginBottom: "1rem", padding: "0.5rem", flex: 1 }}
          />
          <input
            type="text"
            placeholder="Surname"
            value={form.surname}
            onChange={(e) =>
              setForm({ ...form, surname: e.currentTarget.value })
            }
            style={{ marginBottom: "1rem", padding: "0.5rem", flex: 1 }}
          />
        </div>

        <input
          type="text"
          placeholder="Company"
          value={form.company}
          onChange={(e) => setForm({ ...form, company: e.currentTarget.value })}
          style={{ marginBottom: "1rem", padding: "0.5rem", width: "100%" }}
        />

        <input
          type="text"
          placeholder="Serial number"
          value={form.serial_number}
          onChange={(e) =>
            setForm({ ...form, serial_number: e.currentTarget.value })
          }
          style={{ marginBottom: "1rem", padding: "0.5rem", width: "100%" }}
        />

        <input
          type="text"
          placeholder="Role"
          value={form.role}
          onChange={(e) => setForm({ ...form, role: e.currentTarget.value })}
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
