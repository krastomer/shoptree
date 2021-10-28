module.exports = {
  mode : "jit",
  purge: ['./src/*/.{js,jsx,ts,tsx}', './public/index.html'],
  darkMode: false, // or 'media' or 'class'
  theme: {
    fontFamily: {
      body: ['IBM Plex Sans Thai', 'Kanit'], 
    },
    container: {
      padding: {
        DEFAULT: '1rem',
        sm: '2rem',
        lg: '4rem',
        xl: '5rem',
        '2xl': '6rem',
      },
    },
    extend: {
      textColor: ['active'],
      outline: ['hover', 'active'],
    },

  },
  variants: {
    extend: {
      flexWrap: ['hover', 'focus'],
      flex: ['hover', 'focus'],
      alignSelf: ['hover', 'focus'],
      textDecoration: ['focus-visible'],
      ringColor: ['hover', 'active'],
      gridTemplateRows: ['hover', 'focus'],
      gridTemplateColumns: ['hover', 'focus'],
      borderColor: ['active'],
      margin: ['hover', 'focus'],
      gap: ['hover', 'focus'],
      fill: ['hover', 'focus'],
      colors: {
        'brown-theme': '#9D5B53',
      }
    },
  },
  plugins: [
    require('@tailwindcss/forms'),

  ],
}