/**
 * Select website language based on browser language
 * @param currentLang {string}
 * @param langs {Object.<string, string>}
 * @param defaultLang {string}
 */
function selectWebsiteLanguage(currentLang, langs, defaultLang) {
    const key = "initialLangSelected";
    const isInitialLangSelected = localStorage.getItem(key) === "true";
    const browserLang = navigator.language.split("-")[0];
    if (!isInitialLangSelected) {
        localStorage.setItem(key, "true");
        if (browserLang === currentLang) {
            return;
        }
        let languageBrowserUrl = langs[defaultLang];
        const langIsSupported = langs[browserLang] !== undefined;
        if (langIsSupported) {
            languageBrowserUrl = langs[browserLang];
        }
        window.location.replace(languageBrowserUrl);
    }
}

try {
module.exports = { selectWebsiteLanguage };
} catch (e) {}
