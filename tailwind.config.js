/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    "./internal/views/**/*.{go,templ}",
    "./public/js/**/*.js"
  ],
  theme: {
    extend: {},
  },
  plugins: [
    require('@tailwindcss/forms'),
    require('@tailwindcss/typography'),
  ],
}

