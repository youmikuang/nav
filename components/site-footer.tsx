export function SiteFooter() {
  return (
    <div className="relative mx-auto w-full text-sm md:px-6">
      <footer className="block py-4">
        <div className="mx-auto">
          <hr className="border-b-1 mb-4 border-gray-200" />
          <div className="flex flex-wrap items-center justify-center md:justify-between">
            <div className="w-full px-4 md:w-auto">
              <div className="mb-2 text-center md:mb-0 md:text-left">
                Copyright © 2022-GAOHENG | <a
                  href="https://beian.miit.gov.cn/"
                  target="_blank"
                  className="text-blueGray-500 py-1 text-center text-sm font-semibold md:text-left"
                  rel="noreferrer"
                >
                  Copyright © 2023 - present Creative Li WenKai
                </a>
              </div>
            </div>
          </div>
        </div>
      </footer>
    </div>
  )
}
