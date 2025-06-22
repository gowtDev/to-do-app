/** @type {import('next').NextConfig} */
module.exports = {
  basePath: "/todo-ui", // optional, you can remove it if not needed
  async rewrites() {
    return [
      {
        source: "/todo-ui/api/:path*", // frontend call
        destination: "http://localhost:8080/iam/:path*", // backend path includes /iam
      },
    ];
  },
};
