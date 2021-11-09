module.exports = {
  mode: "jit",
  purge: ["./src/**/*.{js,jsx,ts,tsx}", "./public/index.html"],
  darkMode: false, // or 'media' or 'class'
  theme: {
    fontFamily: {
      body: ["IBM Plex Sans Thai", "Kanit"],
    },
    borderWidth: {
      DEFAULT: '1px',
      '0': '0',
      '2': '2px',
     '3': '3px',
      '4': '4px',
     '6': '6px',
     '8': '8px',
    },
    container: {
      padding: {
        DEFAULT: "1rem",
        sm: "2rem",
        lg: "4rem",
        xl: "5rem",
        "2xl": "6rem",
      },
    },
    extend: {
      textColor: ["active"],
      outline: ["hover", "active"],
    },
  },
  variants: {
    extend: {
      borderStyle: ['hover', 'focus'],
      backgroundOpacity: ['active'],
      outline: ['hover', 'active'],
      borderWidth: ['hover', 'focus'],
    },
  },
  plugins: [
    require("@tailwindcss/forms"),
  ],
};
