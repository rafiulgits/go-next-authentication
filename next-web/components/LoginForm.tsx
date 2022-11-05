import { useState } from "react";

interface Props {
  onSubmit: (data: any) => void;
}

export const LoginForm = (props: Props) => {
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const handleSubmit = () => {
    const data = {
      email,
      password,
    };

    props.onSubmit(data);
  };

  return (
    <form
      onSubmit={(e) => {
        e.preventDefault();
        handleSubmit();
      }}
    >
      <div className="mb-3">
        <label className="form-label">Email</label>
        <input
          type="email"
          className="form-control form-control-sm"
          placeholder="Your Email"
          name="email"
          required
          onChange={(e) => setEmail(e.target.value)}
        />
      </div>

      <div className="mb-3">
        <label className="form-label">Password</label>
        <input
          type="password"
          className="form-control form-control-sm"
          placeholder="Password"
          name="password"
          required
          onChange={(e) => setPassword(e.target.value)}
        />
      </div>

      <div className="d-grid gap-2">
        <button
          className="btn btn-info"
          type="submit"
          onSubmit={(e) => e.preventDefault()}
        >
          Login
        </button>
      </div>
    </form>
  );
};
