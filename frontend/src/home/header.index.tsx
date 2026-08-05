const Header = () => {
    return (
        <div className="flex justify-between items-center pt-4 pb-6 container mx-auto text-black-1">
            <div className="text-primary text-2xl font-bold">
                Task Management
            </div>
            <div className="flex items-center gap-4">
                <button className="cursor-pointer rounded-md">switch</button>
                <button className=" p-2 cursor-pointer hover:bg-primary/60 rounded-md">
                    Sign in
                </button>
                <button className="bg-primary text-white p-2 rounded-md cursor-pointer">
                    Get started
                </button>
            </div>
        </div>
    );
};
export default Header;
