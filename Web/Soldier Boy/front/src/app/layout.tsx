import './globals.css'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <head />
      <body className="bg-[#111] text-[#ffe] syne">{children}</body>
    </html>
  )
}
