const { createGlobPatternsForDependencies } = require("@nx/react/tailwind");
const { join } = require("path");
const { fontFamily } = require("tailwindcss/defaultTheme");

/** @type {import('tailwindcss').Config} */
module.exports = {
  content: [
    join(
      __dirname,
      "{src,pages,components,app}/**/*!(*.stories|*.spec).{ts,tsx,html}",
    ),
    ...createGlobPatternsForDependencies(__dirname),
  ],
  darkMode: ["class"],
  prefix: "",
  theme: {
    container: {
      center: true,
      padding: "2rem",
      screens: {
        "2xl": "1400px",
      },
    },
    extend: {
      fontFamily: {
        sans: ["var(--font-noto-sans-thai)", ...fontFamily.sans],
        mono: ["var(--font-mono)", ...fontFamily.mono],
      },
      colors: {
        // ESC
        carmine: {
          50: "color-mix(in srgb, var(--esc-carmine-50) calc(<alpha-value> * 100%), transparent)",
          100: "color-mix(in srgb, var(--esc-carmine-100) calc(<alpha-value> * 100%), transparent)",
          200: "color-mix(in srgb, var(--esc-carmine-200) calc(<alpha-value> * 100%), transparent)",
          300: "color-mix(in srgb, var(--esc-carmine-300) calc(<alpha-value> * 100%), transparent)",
          400: "color-mix(in srgb, var(--esc-carmine-400) calc(<alpha-value> * 100%), transparent)",
          500: "color-mix(in srgb, var(--esc-carmine-500) calc(<alpha-value> * 100%), transparent)",
          600: "color-mix(in srgb, var(--esc-carmine-600) calc(<alpha-value> * 100%), transparent)",
          700: "color-mix(in srgb, var(--esc-carmine-700) calc(<alpha-value> * 100%), transparent)",
          800: "color-mix(in srgb, var(--esc-carmine-800) calc(<alpha-value> * 100%), transparent)",
          900: "color-mix(in srgb, var(--esc-carmine-900) calc(<alpha-value> * 100%), transparent)",
        },

        // Shadcn
        border:
          "color-mix(in srgb, var(--border) calc(<alpha-value> * 100%), transparent)",
        input:
          "color-mix(in srgb, var(--input) calc(<alpha-value> * 100%), transparent)",
        ring: "color-mix(in srgb, var(--ring) calc(<alpha-value> * 100%), transparent)",
        background:
          "color-mix(in srgb, var(--background) calc(<alpha-value> * 100%), transparent)",
        foreground:
          "color-mix(in srgb, var(--foreground) calc(<alpha-value> * 100%), transparent)",
        primary: {
          DEFAULT:
            "color-mix(in srgb, var(--primary) calc(<alpha-value> * 100%), transparent)",
          foreground:
            "color-mix(in srgb, var(--primary-foreground) calc(<alpha-value> * 100%), transparent)",
        },
        secondary: {
          DEFAULT:
            "color-mix(in srgb, var(--secondary) calc(<alpha-value> * 100%), transparent)",
          foreground:
            "color-mix(in srgb, var(--secondary-foreground) calc(<alpha-value> * 100%), transparent)",
        },
        destructive: {
          DEFAULT:
            "color-mix(in srgb, var(--destructive) calc(<alpha-value> * 100%), transparent)",
          foreground:
            "color-mix(in srgb, var(--destructive-foreground) calc(<alpha-value> * 100%), transparent)",
        },
        muted: {
          DEFAULT:
            "color-mix(in srgb, var(--muted) calc(<alpha-value> * 100%), transparent)",
          foreground:
            "color-mix(in srgb, var(--muted-foreground) calc(<alpha-value> * 100%), transparent)",
        },
        accent: {
          DEFAULT:
            "color-mix(in srgb, var(--accent) calc(<alpha-value> * 100%), transparent)",
          foreground:
            "color-mix(in srgb, var(--accent-foreground) calc(<alpha-value> * 100%), transparent)",
        },
        popover: {
          DEFAULT:
            "color-mix(in srgb, var(--popover) calc(<alpha-value> * 100%), transparent)",
          foreground:
            "color-mix(in srgb, var(--popover-foreground) calc(<alpha-value> * 100%), transparent)",
        },
        card: {
          DEFAULT:
            "color-mix(in srgb, var(--card) calc(<alpha-value> * 100%), transparent)",
          foreground:
            "color-mix(in srgb, var(--card-foreground) calc(<alpha-value> * 100%), transparent)",
        },
      },
      borderRadius: {
        lg: "var(--radius)",
        md: "calc(var(--radius) - 2px)",
        sm: "calc(var(--radius) - 4px)",
      },
      keyframes: {
        "accordion-down": {
          from: { height: "0" },
          to: { height: "var(--radix-accordion-content-height)" },
        },
        "accordion-up": {
          from: { height: "var(--radix-accordion-content-height)" },
          to: { height: "0" },
        },
      },
      animation: {
        "accordion-down": "accordion-down 0.2s ease-out",
        "accordion-up": "accordion-up 0.2s ease-out",
      },
    },
  },
  plugins: [require("tailwindcss-animate")],
};
