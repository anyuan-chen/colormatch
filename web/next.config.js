/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  experimental: {
    appDir: true,
  },
  async redirects() {
    return [
      {
        source: "/spotify/callback",
        destination: "/connect/spotify/thanks",
        permanent: true,
      },
    ];
  },
};

module.exports = nextConfig;
