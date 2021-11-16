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
    objectPosition: {
      bottom: 'bottom',
      center: 'center',
      left: 'left',
      right: 'right',
      'right-bottom': 'right bottom',
      'right-top': 'right top',
    },
    minHeight: {
      '0': '0',
      '1/4': '25%',
      '1/2': '50%',
      '3/4': '75%',
      'full': '100%',
     },
  },
  variants: {
    extend: {
      borderStyle: ['hover', 'focus'],
      backgroundOpacity: ['active'],
      outline: ['hover', 'active'],
      borderWidth: ['hover', 'focus'],
      objectPosition: ['hover', 'focus'],
      alignContent: ['hover', 'focus'],
      alignItems: ['hover', 'focus'],
      minHeight: ['hover', 'focus'],
      height: ['hover', 'focus'],
      backgroundPosition: ['hover', 'focus'],
      alignContent: ['hover', 'focus'],
    },
  },
  plugins: [
    require("@tailwindcss/forms"),
  ],
};
